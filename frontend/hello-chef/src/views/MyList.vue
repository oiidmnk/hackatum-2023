<template>
  <!DOCTYPE html>
  <html lang="en">
  <!--Window-->
  <div class="container flex flex-col min-h-screen min-w-full max-h-full">
    <!--Popup-->
    <div v-if="showNewUserPopup" class="modal fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center">
      <!-- The modal -->
      <div id="popupContainer" class="m-auto p-5 border w-1/3 shadow-lg rounded-md bg-white">
        <!-- Modal content -->
        <div class="mt-3 text-center">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Welcome!
          </h3>
          <br />
          <h4>
            As a new user you need to select your
            <p v-if="showHardRequirements">preferences</p>
            <p v-if="showPreferences">Hard Requirements</p>
            so make<br />
            sure we can recommend you the best recipes possible!
          </h4>
          <br />
          <!--Hard requirements-->
          <div class="grid grid-cols-3 gap-4" v-if="showHardRequirements">
            <button v-for="(value, key) in hardRequirements" :key="key" :class="{
              'bg-blue-500': !value,
              'bg-blue-700': value,
              'text-white': true,
              'font-bold': true,
              'py-2': true,
              'px-4': true,
              rounded: true,
              'focus:outline-none': true,
              'focus:shadow-outline': true,
            }" @click="toggleRequirement(key)">
              {{ key.replace("_", " ") }}
            </button>
          </div>
          <!-- Switch to preferences -->
          <div class="items-center px-4 py-3" v-if="showHardRequirements">
            <button id="ok-btn" @click="moveToPreferences"
              class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-auto shadow-sm hover:bg-blue-700 focus:outline-none focus:border-blue-700 focus:ring-2 focus:ring-blue-500">
              Set Preferences
            </button>
          </div>
          <!--Preferences-->
          <div class="grid grid-cols-6 gap-4 overflow-scroll" v-if="showPreferences">
            <button v-for="(value, key) in preferences" :key="key" :class="{
              'bg-blue-500': !value,
              'bg-blue-700': value,
              'text-white': true,
              'font-bold': true,
              'py-2': true,
              'px-4': true,
              rounded: true,
              'focus:outline-none': true,
              'focus:shadow-outline': true,
            }" @click="togglePreference(key)">
              {{ key.replace("_", " ") }}
            </button>
          </div>
          <!-- Close Popup -->
          <div class="items-center px-4 py-3" v-if="showPreferences">
            <button id="ok-btn" @click="closeModal"
              class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-auto shadow-sm hover:bg-blue-700 focus:outline-none focus:border-blue-700 focus:ring-2 focus:ring-blue-500">
              Finished
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- Container for all Dish related Details -->
    <div class="container flex flex-grow min-w-full">
      <!-- List -->
      <div class="container flex-grow w-2/6 bg-green-800 p-4">
        <ul class="space-y-4">
          <li class="" v-for="dish in filteredDishes" :key="dish.Id">
            <div class="container min-w-full flex h-20 bg-gray-400 rounded cursor-pointer" @click="selectDish(dish.Id)">
              <div class="w-20 h-20">
                <img :src="getImageSrc(dish.Image)" alt="Dish Image" class="w-full h-full object-cover" />
              </div>
              <div class="flex-grow flex flex-col justify-center px-2">
                <h2 class="text-lg font-bold">{{ dish.Name }}</h2>
                <div class="flex">
                  <span v-for="(tag, index) in dish.Tags" :key="tag" class="text-sm">
                    {{ tag
                    }}<span v-if="index < dish.Tags.length - 1">, </span>
                  </span>
                </div>
                <!--CookingTime & CookingLevel-->
                <div class="flex space-x-2">
                  <div class="flex items-center justify-center">
                    <p class="mr-1">Time:</p>
                    <!-- Iterate over the range created by the cookingTime -->
                    <font-awesome-icon v-for="n in dish.CookingTime" :key="n" :icon="['far', 'clock']"
                      class="icon-class" />
                  </div>
                  <div class="flex items-center justify-center">
                    <p class="mr-1">Diff:</p>
                    <!-- Iterate over the range created by the cookingLevel -->
                    <font-awesome-icon v-for="n in dish.CookingLevel" :key="n" :icon="['far', 'lemon']"
                      class="icon-class" />
                  </div>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </div>
      <!-- Recipe -->
      <div class="container flex flex-grow bg-gray">
        <!-- Divs between picture and info-->
        <div v-if="showDetails" class="container flex flex-col flex-grow min-w-full">
          <!-- Image -->
          <div v-if="showDetails" class="container h-40 bg-gray-400 min-w-full">
            <div v-if="dishLoading">Loading...</div>
            <div v-else>
              <img :src="getImageSrc(selectedDish.Image)" alt="Dish Image"
                class="w-full h-40 object-cover align-middle" />
            </div>
          </div>
          <!-- Recipe Info-->
          <div class="flex">
            <div class="container flex-grow min-w-full h-full">
              <div v-if="showDetails" class="space-x-4">
                <div v-if="dishLoading">Loading</div>
                <div v-else class="flex text-left">
                  <!-- General Info -->
                  <div class="w-5/6 pr-16 pl-8 py-8">
                    <div class="flex justify-items-center space-x-2">
                      <h1 class="font-bold text-xl">
                        {{ selectedDish.Name }}
                      </h1>
                      <!-- Like button -->
                      <button @click="likeRecipe(selectedDish.Id)"
                        class="p-2 bg-red-500 text-white rounded hover:bg-red-700">
                        Remove
                      </button>
                    </div>
                    <br />
                    <h3 class="text-lg">{{ selectedDish.Description }}</h3>
                    <br />
                    <p>Recipe:</p>
                    <br />
                    <p>{{ selectedDish.CookingInstructions }}</p>
                    <br />
                  </div>
                  <!-- Extra Properties -->
                  <div class="text-right max-w-xs flex-grow flex-wrap bg-green-200 pr-8 py-8">
                    <p>
                      Cooking Time:
                      {{ selectedDish.RecipeProperties.cooking_time }}
                    </p>
                    <font-awesome-icon v-for="n in selectedDish.RecipeProperties.cooking_time" :key="n"
                      :icon="['far', 'clock']" class="icon-class" />
                    <p>
                      Cooking Level:
                      {{ selectedDish.RecipeProperties.cooking_level }}
                    </p>
                    <font-awesome-icon v-for="n in selectedDish.RecipeProperties.cooking_level" :key="n"
                      :icon="['far', 'lemon']" class="icon-class" />
                    <p>Tags:</p>
                    <div class="flex flex-wrap flex-row-reverse">
                      <span v-for="(tag, index) in selectedDish.Tags" :key="tag"
                        class="text-sm text-right justify-items-start">
                        {{ tag }}
                        <span v-if="index < selectedDish.Tags.length - 1">,
                        </span>
                      </span>
                    </div>
                    <p>Properties:</p>
                    <div class="flex-wrap flex flex-row-reverse">
                      <span v-for="(value, key) in selectedDish.Properties" :key="key" class="text-sm">
                        <span v-if="value">
                          {{ key }}
                          <span v-if="notLastItem(key)">, </span>
                        </span>
                      </span>
                    </div>
                  </div>
                </div>
              </div>
              <div></div>
            </div>
          </div>
        </div>
        <div v-else class="flex p-8 h-full">
          <h1 v-if="dishLoading" class="text-2xl font-bold">Please add Items to your Favorites!</h1>
            <h1 v-else class="text-2xl font-bold">Please Select an Item!</h1>
        </div>
      </div>
    </div>
  </div>

  </html>
</template>
  
<script>
export default {
  data() {
    return {
      currentDish: {},
      dishLoading: true,
      filteredDishes: [],
      searchQuery: "",
      showTags: false,
      selectedTags: [],
      allTags: ["asian", "poultry", "spicy", "indian", "healthy", "salads"],
      dishes: [],
      showDetails: false,
      currentUser: {},
      showNewUserPopup: true,
      isDataLoaded: false,
      showHardRequirements: true,
      showPreferences: false,
      userId: "",
      hardRequirements: {
        Vegan: false,
        Vegetarian: false,
        alcohol_free: false,
        mustard_free: false,
        lactose_free: false,
        egg_free: false,
        pork_free: false,
        wheat_free: false,
        soy_free: false,
        mild: false,
        eco_friendly: false,
        gluten_free: false,
        nut_free: false,
      },
      preferences: {},
    };
  },
  watch: {
    showNewUserPopup(newValue) {
      if (newValue === false) {
        this.getRecipes();
      }
    },
  },
  async created() {
    try {
      // Check LocalStorage to see if there is a user set
      if (!localStorage.getItem("userId")) {
        // Wait for the user creation API call to complete before continuing
        await this.createUser();
      } else {
        this.userId = localStorage.getItem("userId"); // Directly use the value from localStorage
        this.showNewUserPopup = false;
      }
      // These methods will only be called after the above await (if present) is resolved
      await this.getTags();
    } catch (error) {
      console.error("An error occurred during the creation process:", error);
    }
  },
  methods: {
    likeRecipe(recipeId) {
      // Method logic to handle liking a recipe
      const apiUrl =
        "http://localhost:8080/like/" + recipeId + "/" + this.userId; // Your API endpoint

      fetch(apiUrl)
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          console.log(recipeId);
          return response.json();
        })
        .then((data) => {
          console.log(data);
          this.getRecipes();
        })
        .catch((error) => {
          this.error =
            "There was a problem with the fetch operation: " + error.message;
        });
    },
    notLastItem(key) {
      console.log("test");
      console.log(this.selectDish.Properties);
      const keys = Object.keys(this.selectedDish.Properties).filter(
        (k) => this.selectedDish.Properties[k]
      );
      return keys.indexOf(key) < keys.length - 1;
    },
    getImageSrc(imagePath) {
      // Use require to resolve the path to the actual image file
      return require(`@/assets/${imagePath}`);
    },
    createUser() {
      // Wrap the fetch call in a Promise
      return new Promise((resolve, reject) => {
        const apiUrl = "http://localhost:8080/user";

        fetch(apiUrl)
          .then((response) => {
            if (!response.ok) {
              throw new Error("Network response was not ok");
            }
            return response.json();
          })
          .then((data) => {
            console.log(data);
            this.userId = data; // Assuming the response contains a userId field
            localStorage.setItem("userId", this.userId);
            console.log("New user created with id: " + this.userId);
            resolve(this.userId); // Resolve the Promise with the userId
          })
          .catch((error) => {
            console.error(
              "There was a problem with the fetch operation:",
              error
            );
            reject(error); // Reject the Promise in case of an error
          });
      });
    },
    getUserId() {
      this.userId = localStorage.getItem("userId");
    },
    toggleRequirement(key) {
      this.hardRequirements[key] = !this.hardRequirements[key];
      console.log(key + " - " + this.hardRequirements[key]);
    },
    moveToPreferences() {
      this.showHardRequirements = false;
      this.showPreferences = true;
    },
    togglePreference(key) {
      this.preferences[key] = !this.preferences[key];
      console.log(key + " - " + this.preferences[key]);
    },
    async closeModal() {
      try {
        // If setHardRequirements and setPreferences are async operations,
        // you should await them to ensure they complete before continuing.
        await this.setHardRequirements(this.hardRequirements);
        await this.setPreferences(this.preferences);

        // This line will only execute after the above await statements are resolved.
        this.showNewUserPopup = false;
      } catch (error) {
        console.error("An error occurred:", error);
        // Handle any errors that might occur during the preferences setting
      }
    },
    setHardRequirements(hardReqs) {
      console.log(hardReqs);
      const apiUrl =
        "http://localhost:8080/user/" + this.userId + "/requirements"; // Your API endpoint

      return fetch(apiUrl, {
        method: "PUT", // or 'PATCH' if you're only partially updating the resource
        headers: {
          "Content-Type": "application/json",
          // Include other headers as needed, like authorization tokens
        },
        body: JSON.stringify({ hardRequirements: hardReqs }), // Send the hardRequirements as JSON
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error(
              "Network response was not ok: " + response.statusText
            );
          }
          console.log(response);
          return response.json();
        })
        .then((data) => {
          // Handle the successful response here
          console.log("Hard requirements updated:", data);
        })
        .catch((error) => {
          // Handle any errors here
          console.error("There was a problem with the update API:", error);
        });
    },
    setPreferences(prefs) {
      console.log(prefs);
      const apiUrl =
        "http://localhost:8080/user/" + this.userId + "/preferences"; // Your API endpoint

      return fetch(apiUrl, {
        method: "PUT", // or 'PATCH' if you're only partially updating the resource
        headers: {
          "Content-Type": "application/json",
          // Include other headers as needed, like authorization tokens
        },
        body: JSON.stringify({ preferences: prefs }), // Send the hardRequirements as JSON
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error(
              "Network response was not ok: " + response.statusText
            );
          }
          return response.json();
        })
        .then((data) => {
          // Handle the successful response here
          console.log("Preferences updated:", data);
        })
        .catch((error) => {
          // Handle any errors here
          console.error("There was a problem with the update API:", error);
        });
    },
    filterList() {
      if (this.searchQuery === "" && this.selectedTags.length === 0) {
        this.filteredDishes = this.dishes;
      }
      if (this.selectedTags.length === 0) {
        this.filteredDishes = this.dishes.filter((dish) =>
          dish.Name.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      } else {
        // General Filter and then filter that list with tags
        this.filteredDishes = this.dishes.filter((dish) =>
          dish.Name.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
        console.log(this.filteredDishes);
        console.log(this.selectedTags);
        this.filteredDishes = this.filteredDishes.filter((dish) =>
          this.selectedTags.some((tag) =>
            dish.Tags.map((dishTag) => dishTag.toLowerCase()).includes(
              tag.toLowerCase()
            )
          )
        );
      }
    },
    // toggleDropdown() {
    //   this.showTags = !this.showTags;
    // },
    // updateSelectedTags(event) {
    //   // Clear the array without losing reactivity
    //   this.selectedTags = [];
    //   if (event.target.value !== "") {
    //     this.selectedTags = Array.from(event.target.options)
    //       .filter((option) => option.selected)
    //       .map((option) => option.value);
    //     this.filterList(this.searchQuery);
    //   } else {
    //     this.filterList(this.searchQuery);
    //   }
    // },
    getTags() {
      const apiUrl = "http://localhost:8080/tags"; // Your API endpoint

      fetch(apiUrl)
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((data) => {
          this.allTags = data; // Set the Tags to the data received from the API
          console.log(this.allTags);
          this.preferences = this.convertTagsToPreferences();
        })
        .catch((error) => {
          console.error("There was a problem with the fetch operation:", error);
        });
    },
    convertTagsToPreferences() {
      const preferences = {};
      this.allTags.forEach((tag) => {
        preferences[tag] = false;
      });
      return preferences;
    },
    getRecipes() {
      const apiUrl = "http://localhost:8080/liked/" + this.userId; // Your API endpoint

      fetch(apiUrl)
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((data) => {
          this.dishes = data; // Set the dishes to the data received from the API
          this.filteredDishes = data;
          console.log(this.filteredDishes);
        })
        .catch((error) => {
          console.error("There was a problem with the fetch operation:", error);
        });
    },
    setDishId(currentId) {
      this.currentDishId = currentId;
    },
    getRecipe(recipeId) {
      console.log("getRecipe Called");
      this.dishLoading = true;
      const apiUrl =
        "http://localhost:8080/recipe/" + recipeId + "/" + this.userId;

      fetch(apiUrl)
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          console.log(recipeId);
          return response.json();
        })
        .then((data) => {
          console.log(recipeId);
          this.selectedDish = data;
          this.dishLoading = false;
          console.log(this.selectedDish);
          console.log(this.dishLoading);
        })
        .catch((error) => {
          this.error =
            "There was a problem with the fetch operation: " + error.message;
          this.dishLoading = false;
        });
    },
    selectDish(recipeId) {
      this.showDetails = true;
      console.log(recipeId);
      this.getRecipe(recipeId);
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