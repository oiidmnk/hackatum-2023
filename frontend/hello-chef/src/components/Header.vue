<template>
    <!-- Header -->
<div
class="container h-20 min-w-full bg-green-800 flex items-center justify-between px-4"
>
<router-link to="/">
  <img src="@/assets/hello_chef_logo.png" alt="Logo" class="h-16" />
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
</template>

<script></script>
<style></style>