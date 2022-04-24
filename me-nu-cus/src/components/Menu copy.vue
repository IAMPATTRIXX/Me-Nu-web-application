<template>
    <v-card class="mx-auto" max-width="1000">
        <!-- โชว์แรงค์ -->
        <v-container fluid v-if="rank_status">
            <div class="text-center">
                <v-dialog v-model="menuinfo" width="1000" :retain-focus="false">
                    <template v-slot:activator="{ on, attrs }">
                        <v-row dense>
                            <v-col v-for="(data,i) in rank_menu" :key="i" :cols="6">
                                <v-card elevation="3" height="auto" width="auto"
                                    @click="getcurrent(data.food[0])" v-bind="attrs" v-on="on">
                                    <!-- <v-card elevation="3" max-height="500" max-width="500"> -->
                                    <v-card-title>
                                        {{data.food[0].name | readmore}}
                                        <v-icon v-show="i==0" color="#FFD700" class="ml-3">fas fa-crown</v-icon>
                                        <v-icon v-show="i==1" color="#C0C0C0" class="ml-3">fas fa-crown</v-icon>
                                        <v-icon v-show="i==2" small color="#CD7F32" class="ml-3">fas fa-crown</v-icon>
                                    </v-card-title>
                                    <v-row no-gutters>
                                        <v-col cols=8>
                                            <v-img class="ml-5" :src="data.food[0].image_name" height="100" width="120">
                                            </v-img>
                                        </v-col>
                                    </v-row>

                                    <v-spacer></v-spacer>

                                    <v-card-text class="mt=5 d-none d-sm-block">
                                        <v-row>
                                            <v-chip class="mr-2 ml-5 mt-2" v-show="data.food[0].type_id==1">
                                                <v-icon left color="green">
                                                    mdi-checkbox-marked-circle
                                                </v-icon>
                                                Normal
                                            </v-chip>
                                            <v-chip class="mr-2 ml-5 mt-2" v-show="data.food[0].type_id==2">
                                                <v-icon left color="green">
                                                    mdi-checkbox-marked-circle
                                                </v-icon>
                                                Vegan
                                            </v-chip>
                                            <v-chip class="mr-2 ml-5 mt-2" v-show="data.food[0].type_id==3">
                                                <v-icon left color="green">
                                                    mdi-checkbox-marked-circle
                                                </v-icon>
                                                Vegetarian
                                            </v-chip>
                                            <v-chip class="mr-2 ml-5 mt-2" v-show="data.food[0].spc_food">
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
                                            <v-icon dense>mdi-currency-usd</v-icon>{{data.food[0].price}} Bath
                                        </v-row>

                                    </v-card-actions>
                                    <v-spacer></v-spacer>
                                </v-card>
                            </v-col>

                        </v-row>

                    </template>
                    <v-card>
                        <v-card-title class="text-h5 grey lighten-2">
                            {{currentname}}
                        </v-card-title>

                        <v-divider></v-divider>

                        <v-card-text class="align-items-end">
                            <v-icon>mdi-food-fork-drink</v-icon>
                            <read-more more-str="read more" :text="currentdesc" less-str="read less" :max-chars="20"></read-more>
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
                            <!-- โชว์หน้าต่าง Ingredients -->
                            <v-row justify="center" class="mt-5">
                                <v-expansion-panels popout>
                                    <v-expansion-panel>
                                        <v-expansion-panel-header>Main Ingredients
                                        </v-expansion-panel-header>
                                        <v-expansion-panel-content>
                                            <v-row dense>
                                                <v-col cols=5>
                                                    <div v-for="(data,i) in currentingre" :key="i">
                                                        {{data.name}}
                                                    </div>
                                                </v-col>
                                                <v-col cols=7>
                                                    <div v-for="(data,i) in currentingre" :key="i">
                                                        {{data.source}}
                                                    </div>
                                                </v-col>
                                            </v-row>
                                        </v-expansion-panel-content>
                                    </v-expansion-panel>
                                </v-expansion-panels>
                            </v-row>
                            <!-- โชว์ Review -->
                            <v-row justify="center" class="mt-5">
                                <v-expansion-panels popout>
                                    <v-expansion-panel>
                                        <v-expansion-panel-header>Review
                                        </v-expansion-panel-header>
                                        <v-expansion-panel-content>
                                            <v-row dense>
                                                <v-col cols=5>
                                                    <div v-for="(data,i) in currentreview" :key="i">
                                                        {{data.createdAt}}
                                                    </div>
                                                </v-col>
                                                <v-col cols=7>
                                                    <div v-for="(data,i) in currentingre" :key="i">
                                                        {{data.reviews}}
                                                    </div>
                                                </v-col>
                                            </v-row>
                                        </v-expansion-panel-content>
                                    </v-expansion-panel>
                                </v-expansion-panels>
                            </v-row>

                        </v-card-text>

                        <v-divider></v-divider>

                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn icon>
                                <v-icon @click="$emit('order',currentname,currentprice)" color="green">
                                    mdi-food-fork-drink
                                </v-icon>
                            </v-btn>
                            <v-btn color="red" text @click="menuinfo = false">
                                Close
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </div>
        </v-container>
        <!-- โชว์ทั่วไป -->
        <v-container fluid v-else>
            <!-- วนนี้ -->
            <div class="text-center">
                <v-dialog v-model="menuinfo" width="1000" :retain-focus="false">
                    <template v-slot:activator="{ on, attrs }">

                        <v-row dense>
                            <v-col v-for="(data,i) in searchmenu" :key="i" :cols="6">
                                <v-card elevation="3" height="auto" width="auto" @click="getcurrent(data)"
                                    v-bind="attrs" v-on="on">
                                    <!-- <v-card elevation="3" max-height="500" max-width="500"> -->
                                    <v-card-title> {{data.name | readmore}}
                                    </v-card-title>
                                    <v-row no-gutters>
                                        <v-col cols=8>
                                            <v-img class="ml-5" :src="data.image_name" height="100" width="120">
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
                                            <v-icon dense>mdi-currency-usd</v-icon>{{data.price}} Bath
                                        </v-row>

                                    </v-card-actions>
                                    <v-spacer></v-spacer>
                                </v-card>
                            </v-col>

                        </v-row>

                    </template>
                    <v-card>
                        <v-card-title class="text-h5 grey lighten-2">
                            {{currentname}}
                        </v-card-title>

                        <v-divider></v-divider>

                        <v-card-text class="align-items-end">
                            <v-icon>mdi-food-fork-drink</v-icon>
                            <read-more more-str="read more" :text="currentdesc" less-str="read less" :max-chars="20"></read-more>
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
                            <!-- โชว์หน้าต่าง Ingredients -->
                            <v-row justify="center" class="mt-5">
                                <v-expansion-panels popout>
                                    <v-expansion-panel>
                                        <v-expansion-panel-header>Main Ingredients
                                        </v-expansion-panel-header>
                                        <v-expansion-panel-content>
                                            <v-row dense>
                                                <v-col cols=5>
                                                    <div v-for="(data,i) in currentingre" :key="i">
                                                        {{data.name}}
                                                    </div>
                                                </v-col>
                                                <v-col cols=7>
                                                    <div v-for="(data,i) in currentingre" :key="i">
                                                        {{data.source}}
                                                    </div>
                                                </v-col>
                                            </v-row>
                                        </v-expansion-panel-content>
                                    </v-expansion-panel>
                                </v-expansion-panels>
                            </v-row>
                            <!-- โชว์ Review -->
                            <v-row justify="center" class="mt-5">
                                <v-expansion-panels popout>
                                    <v-expansion-panel>
                                        <v-expansion-panel-header>Review
                                        </v-expansion-panel-header>
                                        <v-expansion-panel-content v-for="(data,i) in currentreview" :key="i">
                                            <v-row dense>
                                                {{data.createdAt}}
                                            </v-row>
                                            <v-row dense justify="end">
                                                {{data.reviews}}
                                            </v-row>
                                            <!-- <v-row dense>
                                                <v-col cols=3>
                                                    <div v-for="(data,i) in currentreview" :key="i">
                                                        {{data.reviews}}
                                                    </div>
                                                </v-col>
                                                <v-col cols=4>
                                                    <div v-for="(data,i) in currentreview" :key="i">
                                                        {{data.createdAt}}
                                                    </div>
                                                </v-col>
                                            </v-row> -->
                                        </v-expansion-panel-content>
                                    </v-expansion-panel>
                                </v-expansion-panels>
                            </v-row>

                        </v-card-text>

                        <v-divider></v-divider>

                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn icon>
                                <v-icon @click="$emit('order',currentname,currentprice)" color="green">
                                    mdi-food-fork-drink
                                </v-icon>
                            </v-btn>
                            <v-btn color="red" text @click="menuinfo = false">
                                Close
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </div>
        </v-container>
       

    </v-card>
</template>

<script>
    // import axios from 'axios'
    

    export default {
        props: ["menu", "searchmenu", "rank_status", "rank_menu"],
        
        data: () => ({
            list: [],
            currentdata: [],
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
            currentingre: [],
            currentreview: [],
            currentdate:[],
            page: 1

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
                this.currentingre = data.ing_name
                this.currentreview = data.review
                this.slicedate(this.currentreview)
                console.log(this.currentingre)
                // console.log(this.currentreview[0].createdAt)
                // console.log(this.currentreview)
            },
            slicedate(date){
                for(var i=0; i<date.length; i++){
                    // this.currentdate[i] = date[i].createdAt
                    // this.currentdate[i] = (this.currentdate[i].slice(0,10)).concat("  "+(this.currentdate[i].slice(11,19)))
                    this.currentreview[i].createdAt = (this.currentreview[i].createdAt.slice(0,10)).concat("  "+(this.currentreview[i].createdAt.slice(11,19)))
                    // this.currentdate[i] = (this.currentdate[i].slice(0,10)).concat("  "+(this.currentdate[i].slice(11,19)))
                    
                }
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