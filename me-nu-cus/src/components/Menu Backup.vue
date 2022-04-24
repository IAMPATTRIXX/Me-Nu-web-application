<template>
    <v-card class="mx-auto" max-width="1000">
        <v-container fluid>
            <v-row dense>
                <!-- วนนี้ -->
                <v-col v-for="(data,i) in searchmenu" :key="i" :cols="6">
                    <v-card elevation="3" height="auto" width="auto">
                    <!-- <v-card elevation="3" max-height="500" max-width="500"> -->
                        <v-card-title v-text="data.name" > </v-card-title>
                            <v-row no-gutters>
                                <v-col cols=8>
                                    <v-img class="ml-5"
                                        :src="data.image_name"
                                        height="auto" width="auto">
                                    </v-img>
                                </v-col>
                            </v-row>
                            
                        <v-spacer></v-spacer>
                        <v-card-text class="mt=5 d-none d-sm-block">
                            <v-row>
                                <v-chip class="mr-2 ml-5 mt-2" v-show="data.type_id==1">
                                    <v-icon left color="green">
                                        mdi-checkbox-marked-circle
                                    </v-icon>
                                    Normal
                                </v-chip>
                                <v-chip class="mr-2 ml-5 mt-2" v-show="data.type_id==2">
                                    <v-icon left color="green">
                                        mdi-checkbox-marked-circle
                                    </v-icon>
                                    Vegan
                                </v-chip>
                                <v-chip class="mr-2 ml-5 mt-2" v-show="data.type_id==3">
                                    <v-icon left color="green">
                                        mdi-checkbox-marked-circle
                                    </v-icon>
                                    Vegetarian
                                </v-chip>
                                <v-chip class="mr-2 ml-5 mt-2" v-show="data.spc_food">
                                    <v-icon left color="green">
                                        mdi-checkbox-marked-circle
                                    </v-icon>
                                    Promotion!!!
                                </v-chip>
                            </v-row>
                            
                        </v-card-text>
                        <v-spacer></v-spacer>
                        
                        <v-card-actions class="d-flex justify-end">
                            <v-row dense>
                                <v-icon dense class="ml-5">mdi-currency-usd</v-icon>{{data.price}} Bath
                            </v-row>
                            
                                
                            <v-btn icon >
                                <v-icon @click="$emit('order',data)" color="green" >mdi-food-fork-drink</v-icon>
                            </v-btn>
                            
                            <div class="text-center">
                                <v-dialog v-model="menuinfo" width="1000" :retain-focus="false">
                                    <template v-slot:activator="{ on, attrs }">
                                        <v-btn icon v-bind="attrs" v-on="on" class="" @click="getcurrent(data)">
                                                <v-icon>mdi-information</v-icon>
                                        </v-btn>
                                    </template>
                                    <v-card>
                                        <v-card-title class="text-h5 grey lighten-2">
                                            Infomation
                                        </v-card-title>

                                        <v-divider></v-divider>

                                        <v-card-text class="align-items-end">
                                            <v-icon>mdi-food-fork-drink</v-icon>
                                            : {{currentdesc}}
                                        </v-card-text>

                                        <v-card-text class="mt=5">
                                            
                                            <v-row>
                                                <v-chip class="mr-2 ml-5 mt-2" v-show="currenttypeid==1">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Normal
                                                </v-chip>
                                                <v-chip class="mr-2 ml-5 mt-2" v-show="currenttypeid==2">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Vegan
                                                </v-chip>
                                                <v-chip class="mr-2 ml-5 mt-2" v-show="currenttypeid==3">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Vegetarian
                                                </v-chip>
                                                <v-chip class="mr-2 ml-5 mt-2" v-show="currentspe">
                                                    <v-icon left color="yellow">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Promotion!!!
                                                </v-chip>
                                            </v-row>
                                            <v-row class="mt-3 ml-2">
                                                <v-chip class="mr-2  mt-2" v-show="currentseafood">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Seafood
                                                </v-chip>
                                                <v-chip class="mr-2  mt-2" v-show="currentmilk">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Milk
                                                </v-chip>
                                                <v-chip class="mr-2 mt-2" v-show="currentbean">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Bean
                                                </v-chip>
                                                <v-chip class="mr-2 mt-2" v-show="currentegg">
                                                    <v-icon left color="green">
                                                        mdi-checkbox-marked-circle
                                                    </v-icon>
                                                    Egg
                                                </v-chip>
                                            </v-row>

                                        </v-card-text>

                                        <v-divider></v-divider>

                                        <v-card-actions>
                                            <v-spacer></v-spacer>
                                            <v-btn color="red" text @click="menuinfo = false">
                                                Close
                                            </v-btn>
                                        </v-card-actions>
                                    </v-card>
                                </v-dialog>
                            </div>
                        </v-card-actions>
                        <v-spacer></v-spacer>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </v-card>
</template>

<script>
    // import axios from 'axios'

    export default {
        props: ["menu","searchmenu"],
        data: () => ({
            list: [],
            currentdata: [

            ],
            currentNo: 0,
            menuinfo: false,
            currentdesc: "",
            currentname: "",
            currentegg: false,
            currentseafood: false,
            currentbean: false,
            currentmilk: false,
            currentspe: false,
            currentprice: 0,
            currenttypeid: 1,
            
        }),
        computed: {

        },
        methods: {
            getcurrent(data) {
                this.currentname = data.name
                this.currentdesc = data.description
                this.currentegg = data.has_egg
                this.currentseafood = data.has_seafood
                this.currentmilk = data.has_milk
                this.currentbean = data.has_bean
                this.currentprice = data.price
                this.currenttypeid = data.type_id
                console.log(this.currentdesc + this.currentname)
            }
        }

    }
</script>


<style>
    .chip {
        margin-left: 5px;
        margin-top: 2px;
        margin-right: 2px;
        ;
    }
</style>