package main

import (
    "net/http"
	"strconv"

    "github.com/labstack/echo/v4"
)

type Profile struct {
    Name   string
    Health int
    Power  int
    Exp    int
}

func MakeProfile(name string) Profile {
    switch name {
    case "Sasuke":
        return Profile{Name: "Sasuke", Health: 200, Power: 100, Exp: 0}
    case "Goku":
        return Profile{Name: "Goku", Health: 400, Power: 300, Exp: 100}
    case "Naruto":
        return Profile{Name: "Naruto", Health: 150, Power: 200, Exp: 50}
    default:
        return Profile{}
    }
}

func PowerUp(p Profile, multiplier int) Profile {
    p.Health += p.Health * multiplier
    p.Power += p.Power * multiplier
    p.Exp += p.Exp * multiplier
    return p
}

func profileHandler(c echo.Context) error {
    name := c.Param("name")
    profile := MakeProfile(name)
    return c.JSON(http.StatusOK, map[string]interface{}{
        "Name":   profile.Name,
        "Health": profile.Health,
        "Power":  profile.Power,
        "Exp":    profile.Exp,
    })
}

func powerupeHandler(c echo.Context) error {
    name := c.Param("name")
	multiplier, err := strconv.Atoi(c.Param("multiplier"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid multiplier",
        })
    }

    Profile := MakeProfile(name)
    PowerUp := PowerUp(Profile, multiplier)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "Name":   PowerUp.Name,
        "Health": PowerUp.Health,
        "Power":  PowerUp.Power,
        "Exp":    PowerUp.Exp,
    })
}

func main() {
    e := echo.New()

    e.GET("/:name", profileHandler)
	e.GET("/:name/:multiplier", powerupeHandler)

    // Start the server
    e.Start(":8000")
}
