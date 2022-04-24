<template>
  <v-app>
<Header />  
  <v-main>
    <v-container>
      <StaffHeader />
      <div class="mt-5">
        <StaffView/>
      </div>
    </v-container>
  </v-main>
</v-app>
  <!-- <div> -->
    <!-- <StaffHeader /> -->
    <!-- <v-row dense class="mt-5 d-flex flex-row-reverse">
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/add" >
        Add Menu
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/edit" >
        Edit Menu
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff" >
        Main
      </v-btn>
    </v-row> -->
    <!-- <div class="mt-5">
      <Staffview />
    </div>
  </div> -->
</template>

<script>
  import StaffView from '@/components/Staff_OrderView.vue'
  import StaffHeader from '@/components/Staff_Header.vue'
  import Header from '@/components/HeaderView'
  export default {
    components: {
      StaffView,
      Header,
      StaffHeader 
    }
  }
</script>

<style>
  .a {
    float: right;
  }

  .b {
    display: block;
    margin: 20px;
  }
</style>