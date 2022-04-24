<template>
    
    
    
    <v-card class="mx-auto" max-width="1200"  >
        
            <v-row dense>
                <v-col v-for="(orders,index) in order" :key="index" cols="6" >
                    <v-badge overlap color="red">
                        <v-card class="mx-auto" outlined>
                        <v-list-item three-line>
                            <v-list-item-content>
                                <div class="text-overline mb-1">
                                    ID  : {{orders.bill_id}}
                                </div>
                                <v-list-item-title class="text-h6 mb-5">
                                    Table : {{orders.tableNO}} 
                                </v-list-item-title>
                                
                                <v-list-item-subtitle class="text-h5 mb-1">
                                    <div v-for="(orderr,i) in orders.food_name" :key="i">
                                        {{orderr}} 
                                    </div>
                                    <!-- {{showFoodname(orders.food_name)}} -->
                                    <!-- {{orders.food_name}} -->
                                    <!-- {{orders.food_name[0]}}
                                    {{showFoodname(orders.food_name)}} -->
                                </v-list-item-subtitle>
                            </v-list-item-content>
                        </v-list-item>

                        <v-card-actions>
                            <v-btn outlined rounded text right>
                                Done
                            </v-btn>
                        </v-card-actions>
                        </v-card>
                    </v-badge>
                </v-col>
            </v-row>
       
    </v-card>
</template>



<script>
import axios from 'axios'
    export default {
        data: () => ({
            
            menuinfo: false,
            order:[],
            test:[]
        }),

        mounted() {
      axios
        .get('http://localhost:8081/bill/get-order')
        .then(res => {
            this.order = res.data
            this.order = this.order.data.item.order
            // console.log(res.data)
            // console.log('Success')
            // console.log('------------------------------------------------------')
            
            console.log(this.order)
            // this.test = JSON.stringify(this.order[4].food_name)
            // // this.test = JSON.parse(this.test)
            // console.log("JSON test : " + this.test)
        })
        
    },
    methods:{
        showFoodname (text1){
            // for (let i ; i<=n ;i++){
            //     return "Foodname : " + text1[i] + "Quantity : " + text2[i]
            // }
                return text1.lenght
            // return text.length
        }
    },
        
    computed:{
        // parse(){
        //     string.slice(startingindex, endingindex)
        //     return test
        // }
    }


    }
</script>

<style>

</style>