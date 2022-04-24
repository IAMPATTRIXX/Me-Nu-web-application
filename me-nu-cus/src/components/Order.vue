<template>
    <v-dialog v-model="menuinfo" width="1000" :retain-focus="false">
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" class="ma-2" color="primary" dark>
              Order
              <v-icon dark right> mdi-food-fork-drink </v-icon>
            </v-btn>
          </template>
          <v-card>
            <v-card-title class="text-h5 grey lighten-2">
              Infomation
            </v-card-title>

            <v-divider></v-divider>

            <v-card-text class="align-items-end">
              <v-icon>mdi-food-fork-drink</v-icon>
              <font-awesome-icon icon="fa-solid fa-stroopwafel" />



            </v-card-text>

            <v-card-text class="mt=5">
              <v-row>
                <v-chip class="mr-2">
                  <v-icon left color="green">
                    mdi-checkbox-marked-circle
                  </v-icon>
                  Vegan
                </v-chip>
              </v-row>

            </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" text @click="menuinfo = false">
                Order
              </v-btn>
              <v-btn color="primary" text @click="menuinfo = false">
                Close
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
</template>

<script>
export default {
    
}
</script>


<style>
 
</style>