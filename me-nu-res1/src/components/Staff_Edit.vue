<template>
    <v-app>
        <Header />
        <v-main>
            <v-container>
                <div>
                    <StaffHeader />
                    <v-card class="mx-auto mt-5" max-width="1000">
                        <v-container fluid>
                            <v-row dense>
                                <v-col v-for="(menu,index) in menures" :key="index" :cols="4">
                                    <v-card>
                                        <v-row dense>
                                            <v-col cols="9">
                                                <v-card-title>{{menu.name | readmore}}</v-card-title>
                                            </v-col>
                                            <v-col cols="2">
                                                <v-switch v-model="menu.status" flat
                                                    @change="onToggle(menu.status,menu.id)">
                                                </v-switch>

                                            </v-col>
                                        </v-row>

                                        <v-card-actions>
                                            <v-spacer></v-spacer>
                                            
                                            <v-btn color="red" text @click="onDelete(menu.id)">
                                                <v-icon>mdi-trash-can</v-icon>

                                            </v-btn>
                                            <div class="text-center">
                                                <v-dialog v-model="menuinfo" width="1000" :retain-focus="false">
                                                    <template v-slot:activator="{ on, attrs }">
                                                        <v-btn icon v-bind="attrs" v-on="on" @click="getcurrent(menu)">
                                                            <v-icon>mdi-cog-outline</v-icon>
                                                        </v-btn>
                                                    </template>
                                                    <v-card>
                                                        <v-card-title class="text-h5 grey lighten-2">
                                                            Infomation
                                                        </v-card-title>

                                                        <v-divider></v-divider>

                                                        <div>
                                                            <v-form>
                                                                <v-container fluid>
                                                                    <v-row dense class="ml-5 mt-5">
                                                                        <v-col cols="4">
                                                                            <v-text-field dense label=Name type="text"
                                                                                name="menudish" v-model="currentname">
                                                                            </v-text-field>
                                                                        </v-col>
                                                                        
                                                                        <v-col cols="1">
                                                                            <v-text-field dense label=Price name="prize"
                                                                                type="number"
                                                                                v-model.number="currentprice">
                                                                            </v-text-field>
                                                                        </v-col>


                                                                    </v-row>
                                                                    <v-row dense class="ml-5">
                                                                        <v-col cols="6">
                                                                            <v-text-field dense label="Description"
                                                                                type="text" name="desc"
                                                                                v-model="currentdesc">
                                                                            </v-text-field>
                                                                        </v-col>
                                                                    </v-row>
                                                                    <v-row justify="center" class="mt-2">
                                                                        <v-expansion-panels popout>
                                                                            <v-expansion-panel>
                                                                                <v-expansion-panel-header>Review
                                                                                </v-expansion-panel-header>
                                                                                <v-expansion-panel-content
                                                                                    v-for="(data,i) in currentreview"
                                                                                    :key="i">
                                                                                    <v-row dense>
                                                                                        {{data.createdAt}}
                                                                                    </v-row>
                                                                                    <v-row dense justify="end">
                                                                                        {{data.reviews}}
                                                                                    </v-row>
                                                         
                                                                                </v-expansion-panel-content>
                                                                            </v-expansion-panel>
                                                                        </v-expansion-panels>
                                                                    </v-row>
                                                                    <v-row dense>
                                                                        <v-col cols="2" class="checkbox">
                                                                            <v-checkbox label="Seafood?" color="primary"
                                                                                hide-details name="seafood"
                                                                                v-model="currentseafood">
                                                                            </v-checkbox>
                                                                        </v-col>
                                                                        <v-col cols="2">
                                                                            <v-checkbox label="Milk?" color="primary"
                                                                                hide-details name="milk"
                                                                                v-model="currentmilk">
                                                                            </v-checkbox>
                                                                        </v-col>
                                                                        <v-col cols="2">
                                                                            <v-checkbox label="Peanut?" color="primary"
                                                                                hide-details name="peanut"
                                                                                v-model="currentbean">
                                                                            </v-checkbox>
                                                                        </v-col>
                                                                        <v-col cols="2">
                                                                            <v-checkbox label="Egg?" color="primary"
                                                                                hide-details name="egg"
                                                                                v-model="currentegg">
                                                                            </v-checkbox>
                                                                        </v-col>
                                                                        <v-col cols="2">
                                                                            <v-checkbox label="Special Food?"
                                                                                color="primary" hide-details name="spe"
                                                                                v-model="currentspe">
                                                                            </v-checkbox>
                                                                        </v-col>
                                                                    </v-row>
                                                                </v-container>
                                                            </v-form>
                                                        </div>

                                                        <v-divider></v-divider>

                                                        <v-card-actions>
                                                            <v-spacer></v-spacer>
                                                            <!-- <v-btn color="primary" text @click="menuinfo = false">
                                                                Order
                                                            </v-btn> -->

                                                            <v-btn color="primary" text @click="onUpdate">
                                                                Update
                                                            </v-btn>
                                                            <v-btn color="primary" text @click="menuinfo = false">
                                                                Close
                                                            </v-btn>
                                                        </v-card-actions>
                                                    </v-card>
                                                </v-dialog>
                                            </div>
                                        </v-card-actions>
                                    </v-card>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card>
                </div>
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
        data: () => ({
            menuinfo: false,
            menures: [],
            //เรียกมาโชว์
            select: null,
            checkup: false,
            menudish: null,
            price: 0,
            type: null,
            desc: null,
            status: true,
            promo: false,
            checkbox: false,
            seafood: false,
            milk: false,
            peanut: false,
            egg: false,
            //ตัว current ไว้แก้ไข
            currentdesc: "",
            currentname: "",
            currentegg: false,
            currentseafood: false,
            currentbean: false,
            currentmilk: false,
            currentspe: false,
            currentprice: 0,
            currenttypeid: 1,
            currentstatus: false,
            currentid: 0,
            currentreview: [],


        }),

        mounted() {
            // setInterval(() => {
            //    gatdata()
            // },50000)
        },

        created(){
            this.getdata()
        },
        methods: {
            onDelete(id) {
                const fd = {
                    "id": id,
                };
                const test = JSON.stringify(fd)
                console.log('JSON......' + test)
                axios.post(process.env.VUE_APP_BASE_URL+'/food/delete-food', test)
                    .then(res => {
                        console.log(res)
                        this.getdata()
                    })
                // this.getdata()
            },
            //ไว้เก็บค่าตัวปัจจุบัน
            getcurrent(data) {
                this.currentname = data.name
                this.currentdesc = data.description
                this.currentegg = data.has_egg
                this.currentseafood = data.has_seafood
                this.currentmilk = data.has_milk
                this.currentbean = data.has_bean
                this.currentprice = data.price
                this.currenttypeid = data.type_id
                this.currentspe = data.promo
                this.currentstatus = data.status
                this.currentid = data.id
                this.currentreview = data.review
                this.slicedate(this.currentreview)
                console.log(this.currentreview)
            },

            onUpdate() {
                const fd = {
                    "name": this.currentname,
                    "price": this.currentprice,
                    "description": this.currentdesc,
                    "has_milk": this.currentmilk,
                    "has_egg": this.currentegg,
                    "has_bean": this.currentbean,
                    "has_seafood": this.currentseafood,
                    "type_id": this.currenttypeid,
                    "spc_food": this.currentspe,
                    "status": this.currentstatus,
                    "id": this.currentid
                };
                const test = JSON.stringify(fd)
                axios.post(process.env.VUE_APP_BASE_URL+'/food/update-food ', test)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })

            },

            slicedate(date) {
                for (var i = 0; i < date.length; i++) {
                    this.currentreview[i].createdAt = (this.currentreview[i].createdAt.slice(0, 10)).concat("  " + (this
                        .currentreview[i].createdAt.slice(11, 19)))
                }
            },

            onToggle(status, id) {
                const fd = {
                    "id": id,
                    "status": status
                }
                const test = JSON.stringify(fd)
                console.log(test)
                axios.post(process.env.VUE_APP_BASE_URL+'/food/status/toggle', test)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })
            },
            getdata(){
                axios
                .get(process.env.VUE_APP_BASE_URL+'/food/res/menu-food')
                .then(res => {
                    this.menures = res.data.data.item.food
                })
            }
        }


    }
</script>


<style>

</style>