<template>
    <v-card class="mx-auto" max-width="1200">

        <v-row dense >
            <v-col v-for="(orders,index) in order" :key="index" cols="4">
                <v-badge overlap color="red">
                    <v-card class="mx-auto" outlined width="350">

                        <v-card-subtitle class="text-overline mb-1">
                            ID : {{orders.bill_id}}
                        </v-card-subtitle>
                        <v-card-title class="text-h6 mb-5">
                            Table : {{orders.tableNO}}
                        </v-card-title>
                        <v-card-text>
                            <v-row dense>
                                <v-col cols=9>
                                    <div v-for="(orderr,i) in orders.food_name" :key="i">
                                        {{orderr}}
                                    </div>
                                </v-col>
                                <v-col cols=3>
                                    <div v-for="(orderr,i) in orders.quantity" :key="i">
                                        Qty : {{orderr}}
                                    </div>
                                </v-col>
                            </v-row>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn outlined rounded text right @click="OrderStatus(orders.bill_id)">
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
            order: [],
            test: []
        }),

        mounted() {
            // setInterval(() => {
            //     console.log('5555')
            // },3000)
            axios
                .get(process.env.VUE_APP_BASE_URL+'/bill/get-order')
                .then(res => {
                    this.order = res.data
                    this.order = this.order.data.item.order
                    // console.log(res.data)
                    // console.log('Success')
                    // console.log('------------------------------------------------------')

                    // console.log(this.order)
                    // this.test = JSON.stringify(this.order[4].food_name)
                    // // this.test = JSON.parse(this.test)
                    // console.log("JSON test : " + this.test)
                })

        },
        methods: {
            showFoodname(text1) {
                // for (let i ; i<=n ;i++){
                //     return "Foodname : " + text1[i] + "Quantity : " + text2[i]
                // }
                return text1.lenght
                // return text.length
            },
            async OrderStatus(bill_id) {
                const fd = {
                    "id": bill_id,
                    "done": true,
                }
                const fd2 = {
                    "bill_id": bill_id,
                    "is_done": true,
                }
                const test1 = JSON.stringify(fd)
                const test2 = JSON.stringify(fd2)
                console.log(test1)
                console.log(test2)

                axios
                    .post(process.env.VUE_APP_BASE_URL+'/bill/create-bill', test1)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })

                axios
                    .post(process.env.VUE_APP_BASE_URL+'/bill/update-order', test2)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })
            },
            // async BillStatus(){
            //     await this.OrderStatus();
            //     const fd2 ={
            //         "id": bill_id,
            //         "is_done" : true,
            //     }
            //     const test2 = JSON.stringify(fd2)
            //     console.log(test2)
            //     axios.post('http://localhost:8081/bill/update-order',test2)
            //     .then(res => {
            //             console.log(res)
            //         }).catch(error => {
            //             console.log(error)
            //         })
            // }
        },

        computed: {
            // parse(){
            //     string.slice(startingindex, endingindex)
            //     return test
            // }
        }


    }
</script>

<style>

</style>