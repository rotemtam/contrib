// Code generated by entc, DO NOT EDIT.

package blogpost

const (
	// Label holds the string label denoting the blogpost type in the database.
	Label = "blog_post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"

	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"

	// Table holds the table name of the blogpost in the database.
	Table = "blog_posts"
	// AuthorTable is the table the holds the author relation/edge.
	AuthorTable = "blog_posts"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "blog_post_author"
	// CategoriesTable is the table the holds the categories relation/edge. The primary key declared below.
	CategoriesTable = "category_blog_posts"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
)

// Columns holds all SQL columns for blogpost fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldBody,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the BlogPost type.
var ForeignKeys = []string{
	"blog_post_author",
}

var (
	// CategoriesPrimaryKey and CategoriesColumn2 are the table columns denoting the
	// primary key for the categories relation (M2M).
	CategoriesPrimaryKey = []string{"category_id", "blog_post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
