package jsonc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	json := `
  {  //	hello
    "c": 3,"b":3, // jello
    /* SOME
       LIKE
       IT
       HAUT */
    "d\\\"\"e": [ 1, /* 2 */ 3, 4, ],
  }`
	expect := `
  {    	     
    "c": 3,"b":3,         
           
           
         
              
    "d\\\"\"e": [ 1,         3, 4  ] 
  }`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestNewLines(t *testing.T) {
	json := `
  {
    "recipeCuisine": "European
		cuisine",
    /*"nutrition": {
      "@type": "NutritionInformation",
      "calories": "270 calories"
    },*/
    "recipeIngredient": ["200 g Beets (baked or boiled)
      ", "100 g Cherries (fresh or frozen)
      ", "100 g Arugula
      ", "150 g Feta
	  "],
  }`
	expect := `
  {
    "recipeCuisine": "European\n		cuisine",

                    
                                      
                                
        
    "recipeIngredient": ["200 g Beets (baked or boiled)\n      ",
 "100 g Cherries (fresh or frozen)\n      ",
 "100 g Arugula\n      ",
 "150 g Feta\n	  "
] 
  }`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}

func TestInvalidEscapeChar(t *testing.T) {
	json := `
  {
    "description": "This meal has it all\u2014flavorful wild rice. Don\'t have time to dedicate hours.",
  }`
	expect := `
  {
    "description": "This meal has it all\u2014flavorful wild rice. Don't have time to dedicate hours." 
  }`
	assert.Equal(t, expect, string(ToJSON([]byte(json))))
}
