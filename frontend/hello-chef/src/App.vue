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
          <div
            v-if="showTags"
            class="bg-white border rounded shadow-lg z-10"
          >
            <select class="w-full rounded border-0" @change="updateSelectedTags">
              <option value="">none</option>
              <option
                v-for="tag in allTags"
                :key="tag.value"
                :value="tag.value"
              >
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
                @click="selectDish(dish.id)"
              >
                <div class="w-20 h-20">
                  <img
                    src="https://images.unsplash.com/photo-1546069901-ba9599a7e63c?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxleHBsb3JlLWZlZWR8Mnx8fGVufDB8fHx8fA%3D%3D"
                    alt=""
                    class="w-full h-full object-cover"
                  />
                </div>
                <div class="flex-grow flex flex-col justify-center px-2">
                  <h2 class="text-lg font-bold">{{ dish.name }}</h2>
                  <h4 class="text-sm">{{ dish.tags }}</h4>
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
                <h1>{{selectedDish.name}}</h1>
                <br />
                <h3>Tags: {{selectedDish.tags}}</h3>
                <br /><br />
                <p>{{selectedDish.description}}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
</html>

</template>

<script>
export default{
  data() {
    return {
      selectedDish: {},
      filteredDishes: [],
      searchQuery: "",
      showTags: false,
      selectedTags: [],
      allTags: ["asian", "poultry", "spicy", "indian", "healthy", "salads"],
      dishes: [
        {
          id: 1,
          name: "Chicken Teriyaki",
          tags: ["asian", "poultry"],
          description: "Yummy Chicken Teriyaki Recipe",
        },
        {
          id: 2,
          name: "Spice Symphony",
          tags: ["spicy", "indian"],
          description: "Yummy Curry Dish inspired by Indian classics",
        },
        {
          id: 3,
          name: "Green Bowl",
          tags: ["healthy", "salads"],
          description:
            "If you're looking to combine health and taste this is the dish for you",
        },
      ],
      showDetails: false,
    };
  },
  created() {
    this.filteredDishes = this.dishes;
  },
  methods: {
    selectDish(dishId) {
      this.showDetails = true;
      this.selectedDish = this.dishes.find((dish) => dish.id === dishId);
    },
    filterList() {
      if (this.searchQuery === "" && this.selectedTags.length === 0) {
        this.filteredDishes = this.dishes;
      }
      if(
        this.selectedTags.length === 0) {
        this.filteredDishes = this.dishes.filter((dish) =>
        dish.name.toLowerCase().includes(this.searchQuery.toLowerCase()));
      } else {
        // General Filter and then filter that list with tags
        this.filteredDishes = this.dishes.filter((dish) =>
        dish.name.toLowerCase().includes(this.searchQuery.toLowerCase()));
        console.log(this.filteredDishes);
        this.filteredDishes = this.filteredDishes.filter(dish =>
          this.selectedTags.some(tag =>
            dish.tags.map(dishTag => dishTag.toLowerCase()).includes(tag.toLowerCase())
          )
        );
        console.log(this.filteredDishes);
      }
      
    },
    toggleDropdown() {
      this.showTags = !this.showTags;
    },
    updateSelectedTags(event) {
      // Clear the array without losing reactivity
      this.selectedTags = [];
      if(event.target.value !== "") {
        this.selectedTags = Array.from(event.target.options)
                             .filter(option => option.selected)
                             .map(option => option.value);
        this.filterList(this.searchQuery);
      }
      else {
        this.filterList(this.searchQuery);
      }
    },
    getRecipes() {
      const apiUrl = 'http://localhost:8080/api/recipes'; // Replace with the actual API URL
  
      fetch(apiUrl)
        .then(response => {
          // Check if the request was successful
          if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
          }
          return response.json(); // Parse JSON response body
        })
        .then(data => {
          //this.recipes = data; // Set the fetched data to the recipes array
          console.log(data);
        })
        .catch(error => {
          console.error('There was a problem with the fetch operation:', error);
        });
    }
  },
}
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
