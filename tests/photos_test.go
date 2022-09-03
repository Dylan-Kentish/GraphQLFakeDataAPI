package tests

import (
	"fmt"

	"github.com/Dylan-Kentish/GraphQLFakeDataAPI/api"
	"github.com/Dylan-Kentish/GraphQLFakeDataAPI/data"
	"github.com/Dylan-Kentish/GraphQLFakeDataAPI/tests/testData"
	"github.com/graphql-go/graphql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/exp/maps"
)

var _ = Describe("Photos", func() {
	testData := testData.NewTestData()
	testApi := api.NewAPI(testData)

	photoTests := make([]TableEntry, len(testData.Photos))
	for i, photo := range testData.Photos {
		idString := fmt.Sprint(photo.ID)
		photoTests[i] = Entry(idString, photo.ID)
	}

	userTests := make([]TableEntry, len(testData.Users))
	for i, user := range testData.Users {
		idString := fmt.Sprint(user.ID)
		userTests[i] = Entry(idString, user.ID)
	}

	It("Invalid ID", func() {
		// Query
		query := `{photo(id:-1){id,albumid,description}}`
		params := graphql.Params{Schema: testApi.Schema, RequestString: query}
		r := graphql.Do(params)
		Expect(r.Errors).To(BeEmpty())

		result := r.Data.(map[string]interface{})
		var photo data.Photo
		convertTo(result["photo"], &photo)

		Expect(photo).To(Equal(data.Photo{}))
	})

	DescribeTable("Get photo by ID", func(id int) {
		// Query
		query := fmt.Sprintf(`{photo(id:%v){id,albumid,description}}`, id)
		params := graphql.Params{Schema: testApi.Schema, RequestString: query}
		r := graphql.Do(params)
		Expect(r.Errors).To(BeEmpty())

		result := r.Data.(map[string]interface{})
		var photo data.Photo
		convertTo(result["photo"], &photo)

		Expect(photo).To(Equal(testData.Photos[id]))
	}, photoTests)

	DescribeTable("Get photo by albumID", func(albumId int) {
		// Query
		query := fmt.Sprintf(`{photos(albumid:%v){id,albumid,description}}`, albumId)
		params := graphql.Params{Schema: testApi.Schema, RequestString: query}
		r := graphql.Do(params)
		Expect(r.Errors).To(BeEmpty())

		result := r.Data.(map[string]interface{})
		var Photos []data.Photo
		convertTo(result["photos"], &Photos)

		expected := make([]data.Photo, 0)

		for _, photo := range testData.Photos {
			if photo.AlbumID == albumId {
				expected = append(expected, photo)
			}
		}

		Expect(Photos).To(ContainElements(expected))
	}, userTests)

	It("Get all Photos", func() {
		// Query
		query := `{photos{id,albumid,description}}`
		params := graphql.Params{Schema: testApi.Schema, RequestString: query}
		r := graphql.Do(params)
		Expect(r.Errors).To(BeEmpty())

		result := r.Data.(map[string]interface{})
		var Photos []data.Photo
		convertTo(result["photos"], &Photos)

		Expect(Photos).To(ContainElements(maps.Values(testData.Photos)))
	})

	It("Get limited Photos", func() {
		limit := 5
		// Query
		query := fmt.Sprintf(`{photos(limit:%v){id,albumid,description}}`, limit)
		params := graphql.Params{Schema: testApi.Schema, RequestString: query}
		r := graphql.Do(params)
		Expect(r.Errors).To(BeEmpty())

		result := r.Data.(map[string]interface{})
		var Photos []data.Photo
		convertTo(result["photos"], &Photos)

		Expect(Photos).To(HaveLen(limit))
	})

	Context("Bad Schema", func() {
		badQuery := graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"photo": &graphql.Field{
					Type:        testApi.PhotoType,
					Description: "photo by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Description: "id of the photo",
							Type:        graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Wrong type
						return data.User{}, nil
					},
				},
			},
		})

		badSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
			Query: badQuery,
		})

		DescribeTable("Reterns err when resolving fields", func(field string) {
			// Query
			query := fmt.Sprintf(`{photo(id:0){%s}}`, field)
			params := graphql.Params{Schema: badSchema, RequestString: query}
			r := graphql.Do(params)
			Expect(r.Errors).To(HaveLen(1))
			Expect(r.Errors[0].Message).To(Equal("source is not a api.Photo"))
		},
			Entry("id", "id"),
			Entry("albumid", "albumid"),
			Entry("description", "description"))
	})
})