<template>
  <div class="home">
    <user-table-component></user-table-component>
  </div>
</template>

<script>
// @ is an alias to /src
import UserTableComponent from "@/components/UserTableComponent.vue";

export default {
  name: "Home",
  components: {
    UserTableComponent
  }
};
</script>
