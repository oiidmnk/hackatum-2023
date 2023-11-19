<template>
  <!DOCTYPE html>
  <html lang="en">
    <!--Window-->
    <div class="container flex flex-col min-h-screen min-w-full max-h-full">
      <!-- Header -->
      <div
        class="container h-20 min-w-full bg-green-800 flex items-center justify-between px-4"
      >
        <router-link to="/">
          <img src="@/assets/hello_chef_logo.png" alt="Logo" class="h-full" />
        </router-link>
        <nav>
          <ul class="flex">
            <li class="px-5">
              <router-link to="/" class="text-white">Dashboard</router-link>
            </li>
            <li class="px-5">
              <router-link to="/list" class="text-white">My List</router-link>
            </li>
            <li class="px-5">
              <router-link to="/generate" class="text-white"
                >Generate Recipes</router-link
              >
            </li>
          </ul>
        </nav>
      </div>
      <!--Form to generate a Recipe-->
      <div>
        <h3>Welcome to the Beta Recipe Generator</h3>
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

            <div class="mb-4">
              <label
                for="style"
                class="block text-gray-700 text-sm font-bold mb-2"
                >Food Style:</label
              >
              <select
                v-model="formData.style"
                id="style"
                class="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline"
              >
                <option disabled value="">Please select one</option>
                <option v-for="style in stylesList" :key="style" :value="style">
                  {{ style }}
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
    submitForm() {
      const apiUrl = "http://yourapi.com/endpoint"; // Replace with your API endpoint

      fetch(apiUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          // Add any other headers your API needs
        },
        body: JSON.stringify(this.formData),
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((data) => {
          console.log(data);
          // Handle the response from the server
        })
        .catch((error) => {
          console.error("There was a problem with the fetch operation:", error);
        });
    },
  },
};
</script>
  
  <style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>