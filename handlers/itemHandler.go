package handlers

import (
	// "database/sql"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	// "log"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// This implicitly sets the value of rows to be the
		// rows from the items table of the database
		// if there are none, an error will be thrown.
		rows, err := models.Db.Query("SELECT * FROM Items")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close() // This closes the rows once the function is done.

		var items []models.Item
		for rows.Next() {
			var item models.Item
			var listedAtStr string
			err := rows.Scan(&item.ItemID, &item.UserID, &item.Name, &item.Description, &item.Price, &listedAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Parse the ListedAt string into a time.Time
			layout := "2006-01-02 15:04:05" // adjust this layout to match the format of timestamp
			listedAtTime, err := time.Parse(layout, listedAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			item.ListedAt = listedAtTime // assign the time.Time value directly

			items = append(items, item)
		}

		json.NewEncoder(w).Encode(map[string][]models.Item{"items": items})

	case http.MethodPost:

		var userID int

		if cookie, _ := r.Cookie("session_token"); cookie != nil {
			fmt.Println("	Cookie", cookie.Value, "Exists")
			// Verify the cookie value against a database
			err := models.Db.QueryRow("SELECT UserID FROM Sessions WHERE SessionID = ? AND ExpiresAt > NOW()", cookie.Value).Scan(&userID)
			if err != nil {
				if err == sql.ErrNoRows {
					// No session found with the given session token, or the session has expired
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(""))
					return
				}
				// An error occurred while querying the database
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		err := r.ParseMultipartForm(10 << 20) // limit maxMultipartMemory
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the form fields
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Print("	cannot convert to Float64")
		}

		// Create a new Item instance
		item := models.Item{
			Name:        name,
			Description: description,
			Price:       price,
		}

		// Get the uploaded file
		file, header, err := r.FormFile("photo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Generate a unique random name for the file
		randBytes := make([]byte, 16)
		_, err = rand.Read(randBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileName := hex.EncodeToString(randBytes) + "_" + header.Filename

		// Create the file on disk
		filePath := "/Users/jaycm/OneDrive/Documents/wwccwinter2024Capstone/images/item_photos/" + fileName
		newFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// Write the uploaded file's content to the new file
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //THIS
			return
		}

		item.PhotoURL = filePath
		// Close the file
		err = newFile.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Insert the item into the database
		// Insert item details into the items table
		res, err := models.Db.Exec(`
		INSERT INTO items (UserID, Name, Description, Price, ListedAt)
		VALUES (?, ?, ?, ?, ?)
		`, userID, item.Name, item.Description, item.Price, time.Now())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the ID of the newly inserted item
		itemID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the PhotoURL into the ItemPhotos table with the new ItemID
		_, err = models.Db.Exec(`
		INSERT INTO ItemPhotos (ItemID, PhotoURL)
		VALUES (?, ?)
		`, itemID, item.PhotoURL)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Update the name of the image file to prepend the item's ID
		newFilePath := fmt.Sprintf("/Users/jaycm/OneDrive/Documents/wwccwinter2024Capstone/images/item_photos/%d_%s", itemID, fileName)
		err = os.Rename(item.PhotoURL, newFilePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a relative path for the database
		relativePath := fmt.Sprintf("/images/item_photos/%d_%s", itemID, fileName)

		// Update the photo path in the database
		_, err = models.Db.Exec("UPDATE ItemPhotos SET PhotoURL = ? WHERE ItemID = ?", relativePath, itemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Print("user ", item.Name, " created sucessfully!")

		// Redirect to the login page
		http.Redirect(w, r, "/?message=Item+Created+Successfully", http.StatusSeeOther)

	case http.MethodDelete:
		fmt.Println("	Someone's trying to delete an item!")

		// Parse the URL query parameters
		params := r.URL.Query()

		// Get the itemId parameter
		itemId := params.Get("itemId")
		fmt.Println("	", itemId)

		_, err := models.Db.Exec("DELETE FROM itemphotos WHERE ItemID = ?", itemId)
		if err != nil {
			fmt.Println("	ERROR: Unable to delete from itemphotos")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = models.Db.Exec("DELETE FROM Items WHERE ItemID = ?", itemId)
		if err != nil {
			fmt.Println("	ERROR: Unable to delete from Items")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/?message=Item+Deleted", http.StatusSeeOther)
		fmt.Println("	Item Deleted")
	case http.MethodPut:
		var itemToUpdate models.Item
		err := json.NewDecoder(r.Body).Decode(&itemToUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = models.Db.Exec("UPDATE Items SET Name = ?, Description = ?, Price = ? WHERE ID = ?",
			itemToUpdate.Name, itemToUpdate.Description, itemToUpdate.Price, itemToUpdate.ItemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed) // for all other http methods
	}
}
