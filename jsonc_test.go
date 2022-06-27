package jsonc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	json := `{}`
	expect := `{}`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestComments(t *testing.T) {
	json := `{  //	hello
    "c": 3,"b":3, // jello
    /* SOME
       LIKE
       IT
       HAUT */
    "d\\\"\"e": [ 1, /* 2 */ 3, 4, ],
  }`

	expect := `{    	     
    "c": 3,"b":3,         
           
           
         
              
    "d\\\"\"e": [ 1,         3, 4  ] 
  }`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestNewLines(t *testing.T) {
	json := `{
    "recipeCuisine": "European
		cuisine",
    /*"nutrition": {
      "calories": "270 calories"
    },*/
    "recipeIngredient": ["200 g Beets (baked or boiled)
      ", "150 g Feta
	  "],
  }`
	expect := `{
    "recipeCuisine": "European\n\t\tcuisine",

                    
                                
        
    "recipeIngredient": ["200 g Beets (baked or boiled)\n      ",
 "150 g Feta\n\t  "
] 
  }`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestInvalidEscapeChar(t *testing.T) {
	json := `{
    "description": "This meal has it all\u2014flavorful wild rice. Don\'t have time to dedicate hours.",
  }`
	expect := `{
    "description": "This meal has it all\u2014flavorful wild rice. Don't have time to dedicate hours." 
  }`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestMissingCommas(t *testing.T) {
	json := `{
	"recipeIngredient": [					"400 г Мука 
		", "1 ч.л. Соль 
		", "1 ч.л. Сахар"										"450 г Сулугуни (или имеретинский сыр)
		", "по вкусу Соль и перец"],
	}`
	expect := `{
	"recipeIngredient": [					"400 г Мука \n\t\t",
 "1 ч.л. Соль \n\t\t",
 "1 ч.л. Сахар",										"450 г Сулугуни (или имеретинский сыр)\n\t\t",
 "по вкусу Соль и перец"] 
	}`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestMissingEscaping(t *testing.T) {
	json := `{
	"recipeInstructions": [
	  {
	  "text": "Именно такой режим я использовал в своей <a rel="nofollow" href="https://www.panasonic.com/ua/consumer/kitchen-appliances/steam-ovens/nu-sc300bzpe.html" target="_blank">печи Panasonic</a>.",
	  },]
}`
	expect := `{
	"recipeInstructions": [
	  {
	  "text": "Именно такой режим я использовал в своей <a rel=\"nofollow\" href=\"https://www.panasonic.com/ua/consumer/kitchen-appliances/steam-ovens/nu-sc300bzpe.html\" target=\"_blank\">печи Panasonic</a>." 
	  } ]
}`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}
