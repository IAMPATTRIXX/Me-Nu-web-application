<template>
    <v-row dense class="mt-5 d-flex flex-row-reverse">

      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff/ads" >
        Ads
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff/bill" >
        Bill
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff/qr" >
        Qr Gen
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff/rank" >
        Rank
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff/add" >
        Add Menu
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/staff/edit" >
        Menu
      </v-btn>
      <v-btn outlined rounded text color="grey darken-2" dark router-link to="/" >
        Main
      </v-btn>
    </v-row>
</template>

<script>
export default {
    
}
</script>