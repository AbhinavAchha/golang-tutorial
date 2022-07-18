package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AbhinavAchha/golang-tutorial/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const controllerString = "mongodb+srv://Abhinav:abhinav@cluster0.xhvzlq0.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	var clientOption = options.Client().ApplyURI(controllerString)

	var client, e = mongo.Connect(context.TODO(), clientOption)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection created!")
}

func insertOneMovie(movie model.Netflix) {
	var ok, e = collection.InsertOne(context.Background(), movie)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(&ok)

}

func updateOneMovie(movieId string) {
	var id, _ = primitive.ObjectIDFromHex(movieId)
	var filter = bson.M{"_id": id}
	var update = bson.M{"$set": bson.M{"watched": true}}

	var result, e = collection.UpdateOne(context.Background(), filter, update)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(result.ModifiedCount)
}

func deleteOneMovie(movidId string) {
	var id, _ = primitive.ObjectIDFromHex(movidId)
	var filter = bson.M{"_id": id}
	var result, e = collection.DeleteOne(context.Background(), filter)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(result.DeletedCount)
}

func deleteAllMovie() int64 {
	var result, e = collection.DeleteMany(context.Background(), bson.M{})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(result)
	return result.DeletedCount
}

func getAllMovies() []primitive.M {
	// var result, e = collection.Find(context.Background(), bson.M{})
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// fmt.Println(result)

	var cursor, e = collection.Find(context.Background(), bson.M{})
	if e != nil {
		log.Fatal(e)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		e = cursor.Decode(&movie)
		if e != nil {
			log.Fatal(e)
		}
		movies = append(movies, movie)
	}

	return movies

}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	var allMovies = getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var movieId = mux.Vars(r)["id"]
	println(movieId)
	updateOneMovie(movieId)
	json.NewEncoder(w).Encode(movieId)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	var movieId = mux.Vars(r)["id"]
	deleteOneMovie(movieId)
	json.NewEncoder(w).Encode(movieId)

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	var count = deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
