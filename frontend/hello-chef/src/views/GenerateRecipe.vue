<template>
  <!DOCTYPE html>
  <html lang="en">
    <!--Window-->
    <div class="container flex flex-col min-h-screen min-w-full max-h-full">
      <!--Form to generate a Recipe-->
      <div class="p-16">
        <h3>Welcome to the Beta Version of the Recipe Generator</h3>
        <p>
          Input a series of ingredients and select a tag, your allergens will be
          taken into consideration here
        </p>
        <br />
        <div class="container mx-auto p-4">
          <form @submit.prevent="submitForm">
            <div class="mb-4">
              <label
                for="ingredients"
                class="block text-gray-700 text-sm font-bold mb-2"
                >Ingredients:</label
              >
              <select
                v-model="formData.ingredients"
                id="ingredients"
                class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline"
              >
                <option disabled value="">Please select one</option>
                <option
                  v-for="ingredient in ingredientsList"
                  :key="ingredient"
                  :value="ingredient"
                >
                  {{ ingredient }}
                </option>
              </select>
            </div>

            <button
              type="submit"
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            >
              Submit
            </button>
          </form>
        </div>
      </div>
    </div>
  </html>
</template>
  
  <script>
export default {
  data() {
    return {
      formData: {
        ingredients: "",
        style: "",
      },
      ingredientsList: ["Tomato", "Cheese", "Basil", "Chicken"],
      stylesList: ["Italian", "Mexican", "Indian", "Chinese"],
    };
  },
  methods: {
    submitRecipeForm() {
      const promptString = `Hello, please generate a recipe containing the following ingredients [${this.ingredients.join(
        ", "
      )}], in the style of ${
        this.style
      }. Avoid the following: [${this.avoidances.join(
        ", "
      )}] due to an allergy of the person that will be eating the dish.`;

      // Call your API endpoint and pass the prompt string
      const apiUrl = "/api/generate-recipe"; // Your server endpoint that calls OpenAI
      fetch(apiUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          // Include other headers as needed
        },
        body: JSON.stringify({ prompt: promptString }),
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((data) => {
          console.log(data);
          this.promptResponse = data;
          // Handle the response here, e.g., display the generated recipe
        })
        .catch((error) => {
          console.error("There was a problem with the fetch operation:", error);
        });
    },
  },
};
</script>
export default { 
    data() {
        return {
          promptResponse: "",
        };
      },
}

  <style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>