<template>
    <v-app>
        <Header />
        <v-main>
            <v-container>
                <div>
                    <StaffHeader />
                </div>
            </v-container>
            <v-container>
                <v-row dense>
                    <p class="npaid">Haven't PAID</p>
                </v-row>
                <v-row dense>
                    <v-col v-for="(data,index) in npaid" :key="index" cols="4">
                        <v-card class="mx-auto" outlined width="350">

                            <v-card-subtitle class="text-overline mb-1">
                                ID : {{data.id}}
                            </v-card-subtitle>


                            <v-card-title class="text-h6 mb-5">
                                Table : {{data.table_no}}
                            </v-card-title>
                            <v-hover v-slot="{ hover }" close-delay="300">
                                <v-card-text>
                                    <v-expand-transition>
                                        <div v-if="hover">
                                            <v-row dense>
                                                <v-col cols=7>
                                                    <div v-for="(dataa,i) in data.order_bill" :key="i">
                                                        {{dataa.food_name}}
                                                    </div>
                                                </v-col>
                                                <v-col cols=3>
                                                    <div v-for="(dataa,i) in data.order_bill" :key="i">
                                                        Qty : {{dataa.quantity}}
                                                    </div>
                                                </v-col>
                                            </v-row>
                                        </div>
                                    </v-expand-transition>
                                </v-card-text>
                            </v-hover>
                            <v-card-actions class="d-flex justify-end">
                                <h3 class="mr=5"> Pricetotal : {{amount[index]}} $</h3>
                                <v-btn x-large icon class="pa-3" @click="checkpaid(data.id)">
                                    <v-icon >fas fa-hand-holding-dollar</v-icon>
                                </v-btn>
                            </v-card-actions>
                        </v-card>

                        <!-- <v-card-text>
                                <v-row dense>
                                    <v-col cols=7>
                                        <div v-for="(dataa,i) in data.order_bill" :key="i">
                                            {{dataa.food_name}} Qty : {{dataa.quantity}}
                                        </div>
                                    </v-col>
                                    <v-col cols=3>
                                        <div v-for="(dataa,i) in data.order_bill" :key="i">
                                            Qty : {{dataa.quantity}}
                                        </div>
                                    </v-col>
                                </v-row>
                            </v-card-text> -->
                    </v-col>
                </v-row>
            </v-container>
            <v-container>
                <v-row dense>
                    <p class="paid">PAID</p>
                </v-row>
                <v-row dense>
                    <v-col v-for="(data,index) in paid" :key="index" cols="4">
                        <v-card class="mx-auto" outlined width="350">

                            <v-card-subtitle class="text-overline mb-1">
                                ID : {{data.id}}
                            </v-card-subtitle>


                            <v-card-title class="text-h6 mb-5">
                                Table : {{data.table_no}}
                            </v-card-title>
                            <v-hover v-slot="{ hover }" close-delay="300">
                                <v-card-text>
                                    <v-expand-transition>
                                        <div v-if="hover">
                                            <v-row dense>
                                                <v-col cols=7>
                                                    <div v-for="(dataa,i) in data.order_bill" :key="i">
                                                        {{dataa.food_name}}
                                                    </div>
                                                </v-col>
                                                <v-col cols=3>
                                                    <div v-for="(dataa,i) in data.order_bill" :key="i">
                                                        Qty : {{dataa.quantity}}
                                                    </div>
                                                </v-col>
                                            </v-row>
                                        </div>
                                    </v-expand-transition>
                                </v-card-text>
                            </v-hover>
                            <v-card-actions class="d-flex justify-end">
                                <v-icon x-large class="pa-3">fas fa-hand-holding-dollar</v-icon>
                            </v-card-actions>
                        </v-card>

                        <!-- <v-card-text>
                                <v-row dense>
                                    <v-col cols=7>
                                        <div v-for="(dataa,i) in data.order_bill" :key="i">
                                            {{dataa.food_name}} Qty : {{dataa.quantity}}
                                        </div>
                                    </v-col>
                                    <v-col cols=3>
                                        <div v-for="(dataa,i) in data.order_bill" :key="i">
                                            Qty : {{dataa.quantity}}
                                        </div>
                                    </v-col>
                                </v-row>
                            </v-card-text> -->
                    </v-col>
                </v-row>
            </v-container>


        </v-main>
    </v-app>
</template>

<script>
    import axios from 'axios'
    import StaffHeader from '@/components/Staff_Header.vue'
    import Header from '@/components/HeaderView'

    export default {
        components: {
            StaffHeader,
            Header
        },
        data() {
            return {
                paid: [],
                npaid: [],
                amount : [],
                sum :0
                
            }
        },
        mounted() {
        },
        created(){
            this.getdata()
        },
        methods: {
            checkpaid(id) {
                console.log(id)
                const fd = {
                    "id": id,
                    "is_paid": true
                }
                const send = JSON.stringify(fd)
                axios.post(process.env.VUE_APP_BASE_URL+'/bill/create-bill', send)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })
            },
            getdata(){
                
                axios
                .get(process.env.VUE_APP_BASE_URL+'/bill/get-bill')
                .then(res => {
                    this.npaid = res.data.data.item.bill_npaid
                    console.log(this.npaid)
                    // console.log(this.npaid[0].order_bill[0].price)
                    for(let i =0 ;i<this.npaid.length;i++){
                        for(let j =0;j<this.npaid[i].order_bill.length;j++){
                            this.sum += this.npaid[i].order_bill[j].price
                            this.amount[i] = this.sum
                        }
                        this.sum = 0
                    }
                    
                    // console.log('Success')
                    // console.log('------------------------------------------------------')
                    // this.bill = this.bill.data.item.ranking
                    // console.log(this.bill)
                    // console.log(this.menu)
                })
            }
        }

    }
</script>


<style>
    .npaid {
        font-size: 50px;
        color: red;
    }

    .paid {
        font-size: 50px;
        color: green;
    }
</style>