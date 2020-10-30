<template>
  <div class="home">
    <company-table-component></company-table-component>
  </div>
</template>

<script>
// @ is an alias to /src
import CompanyTableComponent from "@/components/CompanyTableComponent.vue";

export default {
  name: "Companies",
  components: {
    CompanyTableComponent
  }
};
</script>
