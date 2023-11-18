<template>
  <!DOCTYPE html>
  <html lang="en">
    <!-- Header -->
    <div class="container h-20 min-w-full bg-green-800">
      <nav>
        <ul class="flex-1 text-center">
          <li class="list-none inline-block px-5"><p>test</p></li>
          <li class="list-none inline-block px-5"><p>test</p></li>
        </ul>
      </nav>
    </div>
    <!--Popup-->
    <div
      v-if="showNewUserPopup"
      class="modal fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center"
    >
      <!-- The modal -->
      <div
        id="popupContainer"
        class="m-auto p-5 border w-1/3 shadow-lg rounded-md bg-white"
      >
        <!-- Modal content -->
        <div class="mt-3 text-center">
          <h3 class="text-lg leading-6 font-medium text-gray-900">Welcome!</h3>
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
            <button
              v-for="(value, key) in hardRequirements"
              :key="key"
              :class="{
                'bg-blue-500': !value,
                'bg-blue-700': value,
                'text-white': true,
                'font-bold': true,
                'py-2': true,
                'px-4': true,
                rounded: true,
                'focus:outline-none': true,
                'focus:shadow-outline': true,
              }"
              @click="toggleRequirement(key)"
            >
              {{ key.replace("_", " ") }}
            </button>
          </div>
          <!-- Switch to preferences -->
          <div class="items-center px-4 py-3" v-if="showHardRequirements">
            <button
              id="ok-btn"
              @click="moveToPreferences"
              class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-auto shadow-sm hover:bg-blue-700 focus:outline-none focus:border-blue-700 focus:ring-2 focus:ring-blue-500"
            >
              Set Preferences
            </button>
          </div>
          <!--Preferences-->
          <div class="grid grid-cols-3 gap-4" v-if="showPreferences">
            <button
              v-for="(value, key) in preferences"
              :key="key"
              :class="{
                'bg-blue-500': !value,
                'bg-blue-700': value,
                'text-white': true,
                'font-bold': true,
                'py-2': true,
                'px-4': true,
                rounded: true,
                'focus:outline-none': true,
                'focus:shadow-outline': true,
              }"
              @click="toggleRequirement(key)"
            >
              {{ key.replace("_", " ") }}
            </button>
          </div>
          <!-- Close Popup -->
          <div class="items-center px-4 py-3" v-if="showPreferences">
            <button
              id="ok-btn"
              @click="closeModal"
              class="px-4 py-2 bg-blue-500 text-white text-base font-medium rounded-md w-auto shadow-sm hover:bg-blue-700 focus:outline-none focus:border-blue-700 focus:ring-2 focus:ring-blue-500"
            >
              Finished
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- Search -->
    <div class="container min-w-full p-4 space-y-1">
      <!--Search Bar-->
      <div class="min-w-full flex space-x-1">
        <input
          type="search"
          name="search"
          id="search"
          placeholder="Search..."
          class="w-2/6 h-10 px-5 pr-10 rounded text-sm focus:outline-none border-2 border-gray-300 focus:border-green-500"
          @input="filterList"
          v-model="searchQuery"
        />
        <button
          class="h-10 px-3 text-sm rounded border-2 border-gray-300 focus:outline-none focus:border-green-500"
          @click="toggleDropdown"
        >
          Filter
        </button>
        <div v-if="showTags" class="bg-white border rounded shadow-lg z-10">
          <select class="w-full rounded border-0" @change="updateSelectedTags">
            <option value="">none</option>
            <option v-for="tag in allTags" :key="tag.value" :value="tag.value">
              {{ tag }}
            </option>
          </select>
        </div>
      </div>
      <!--tags-->
      <div class="flex min-w-full space-x-2">
        <div></div>
        <!-- <div class="tags rounded" v-for="tag in selectedTags">
            {{ tag }}
          </div> -->
      </div>
    </div>
    <!-- Container for all Dish related Details -->
    <div class="container flex flex-grow min-w-full">
      <!-- List -->
      <div class="container flex-grow w-2/6 bg-green-800 p-4">
        <button @click="getRecipes">get recipes</button>
        <ul class="space-y-4">
          <li class="" v-for="dish in filteredDishes" :key="dish.id">
            <div
              class="container min-w-full flex h-20 bg-gray-400 rounded cursor-pointer"
              @click="selectDish(dish.Id)"
            >
              <div class="w-20 h-20">
                <img
                  :src="'localhost:8080/' + dish.Image"
                  alt=""
                  class="w-full h-full object-cover"
                />
              </div>
              <div class="flex-grow flex flex-col justify-center px-2">
                <h2 class="text-lg font-bold">{{ dish.Name }}</h2>
                <h4 class="text-sm">{{ dish.Tags }}</h4>
              </div>
            </div>
          </li>
        </ul>
      </div>
      <!-- Recipe -->
      <div class="container flex flex-grow bg-gray">
        <!-- Divs between picture and info-->
        <div class="container flex flex-col flex-grow min-w-full">
          <!-- Image -->
          <div class="container h-40 bg-gray-400 min-w-full"></div>
          <!-- Recipe Info-->
          <div class="container flex-grow min-w-full h-full">
            <div v-if="showDetails">
              <div v-if="dishLoading">Loading</div>
              <div v-else>
                <h1>{{ selectedDish.Id }}</h1>
                <br />
                <h3>Tags: {{ selectedDish.Tags }}</h3>
                <p>{{ selectedDish.Description }}</p>
                <br /><br />
              </div>
            </div>
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
      showHardRequirements: true,
      showPreferences: false,
      userId: "",
      hardRequirements: {
        vegan: false,
        vegetarian: false,
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
      preferences: {
        asian: false,
        indian: false,
        italian: false,
        burgers: false,
        thai: false,
        vietnamese: false,
      },
    };
  },
  created() {
    this.getRecipes();
    // check LocalStorage to see if there is a user set
    if (!localStorage.getItem("userId")) {
      // Call the user creation API if not
      this.createUser();
    } else {
      this.showNewUserPopup = false;
      this.userId = this.getUserId();
    }
  },
  methods: {
    createUser() {
      const apiUrl = "http://localhost:8080/user"; // Your API endpoint

      fetch(apiUrl)
        .then((response) => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();
        })
        .then((data) => {
          this.userId = data; // Set the dishes to the data received from the API
          localStorage.setItem('userId', this.userId);
          console.log("New user created with id: " + this.userId);
        })
        .catch((error) => {
          console.error("There was a problem with the fetch operation:", error);
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
    closeModal() {
      this.showNewUserPopup = false;
      this.setHardRequirements(this.hardRequirements);
      this.setPreferences(this.preferences);
    },
    setHardRequirements(hardReqs) {
      const apiUrl = "http://localhost:8080/user/:" + this.userId + "/requirements"; // Your API endpoint

      fetch(apiUrl, {
        method: 'PUT', // or 'PATCH' if you're only partially updating the resource
        headers: {
          'Content-Type': 'application/json',
          // Include other headers as needed, like authorization tokens
        },
        body: JSON.stringify({ hardRequirements: hardReqs }) // Send the hardRequirements as JSON
      })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok: ' + response.statusText);
        }
        return response.json();
      })
      .then(data => {
        // Handle the successful response here
        console.log('Hard requirements updated:', data);
      })
      .catch(error => {
        // Handle any errors here
        console.error('There was a problem with the update API:', error);
      });
    },
    setPreferences(prefs) {
      const apiUrl = "http://localhost:8080/user/:" + this.userId + "/preferences"; // Your API endpoint

      fetch(apiUrl, {
        method: 'PUT', // or 'PATCH' if you're only partially updating the resource
        headers: {
          'Content-Type': 'application/json',
          // Include other headers as needed, like authorization tokens
        },
        body: JSON.stringify({ preferences: prefs }) // Send the hardRequirements as JSON
      })
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok: ' + response.statusText);
        }
        return response.json();
      })
      .then(data => {
        // Handle the successful response here
        console.log('Preferences updated:', data);
      })
      .catch(error => {
        // Handle any errors here
        console.error('There was a problem with the update API:', error);
      });
    },
    filterList() {
      if (this.searchQuery === "" && this.selectedTags.length === 0) {
        this.filteredDishes = this.dishes;
      }
      if (this.selectedTags.length === 0) {
        this.filteredDishes = this.dishes.filter((dish) =>
          dish.name.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
      } else {
        // General Filter and then filter that list with tags
        this.filteredDishes = this.dishes.filter((dish) =>
          dish.name.toLowerCase().includes(this.searchQuery.toLowerCase())
        );
        console.log(this.filteredDishes);
        this.filteredDishes = this.filteredDishes.filter((dish) =>
          this.selectedTags.some((tag) =>
            dish.tags
              .map((dishTag) => dishTag.toLowerCase())
              .includes(tag.toLowerCase())
          )
        );
      }
    },
    toggleDropdown() {
      this.showTags = !this.showTags;
    },
    updateSelectedTags(event) {
      // Clear the array without losing reactivity
      this.selectedTags = [];
      if (event.target.value !== "") {
        this.selectedTags = Array.from(event.target.options)
          .filter((option) => option.selected)
          .map((option) => option.value);
        this.filterList(this.searchQuery);
      } else {
        this.filterList(this.searchQuery);
      }
    },
    getRecipes() {
      const apiUrl = "http://localhost:8080/recipes"; // Your API endpoint

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
      const apiUrl = "http://localhost:8080/recipe/" + recipeId;

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
  margin-top: 60px;
}
</style>
