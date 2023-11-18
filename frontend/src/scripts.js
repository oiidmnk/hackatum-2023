const app = Vue.createApp({
    data() {
      return {
        selectedDish: {

        },
        dishes: [
          {
            id: 1,
            name: 'Chicken Teriyaki',
            tags: ["asian, poultry"],
            description: "Yummy Chicken Teriyaki Recipe",
          },
          {
            id: 2,
            name: 'Spice Symphony',
            tags: ["spicy, indian"],
            description: "Yummy Curry Dish inspired by Indian classics",
          },
          {
            id: 3,
            name: 'Green Bowl',
            tags: ["healthy, salads"],
            description: "If you're looking to combine health and taste this is the dish for you",
          },
        ],
        showDetails:false,
      };
    },
    methods: {
      selectDish(dishId) {
        this.showDetails=true;
        this.selectedDish = this.dishes.find(dish => dish.id === dishId);
    }
    },
  });
  
  app.mount("#window");
  