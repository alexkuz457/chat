package user

mport (
	"net/http"
	"os/user"

	"github.com/labstack/echo"
)

type Request struct {
	Username string `json:"name"`
}

type Response struct {
	User *user.User
}

func AddUser(c echo.Context) (err error) {
	u := new(Request)
	if err = c.Bind(u); err != nil {
		return
	}
	//db := Database()
	//db.Save(u)
	return c.JSON(http.StatusOK, u)
}
