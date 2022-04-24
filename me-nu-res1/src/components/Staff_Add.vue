<template>
<v-app>
    <Header />
    <v-main>
        <v-container>
            <div>
                <StaffHeader />
                <v-card max-width="1000" class="mx-auto mt-5">
                    <v-form>
                        <v-container fluid>
                            <v-row dense>
                                <v-col cols="4" class="Name">
                                    <v-text-field  dense label="Input Menu"  type="text" name="menudish" v-model="menudish">
                                    </v-text-field>
                                </v-col>
                                
                                <v-col cols="4" class="Name">
                                    <v-text-field  dense label="Input Price" name="prize" type="number" v-model.number="price">
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <v-row dense>
                                <v-col cols="6" class="Name">
                                    <v-text-field dense label="Input Description" type="text" name="desc" v-model="desc">
                                    </v-text-field>
                                </v-col>
                            </v-row>
                            <!-- <v-row dense v-for="index in ing.length+1" :key="index"> -->
                            <v-row dense >
                                <!-- <div v-for ="(item,index) in ing" :key="index"> -->
                                <v-col cols="4" class="Name">
                                    <v-text-field dense label="Input Ingredients" type="text" name="ing_name" v-model="ing_name" >
                                    </v-text-field>
                                </v-col>
                                <v-col cols="4" class="Name">
                                    <v-text-field dense label="Input Ingredients source" type="text" name="ing_source" v-model="ing_source">
                                    </v-text-field>
                                </v-col>
                                <v-col cols="1" class ="mt-5 ml-2" >
                                    <v-btn icon elevation="2" @click="onAddIng(ing_name,ing_source)">
                                        <v-icon>mdi-plus-box</v-icon>
                                    </v-btn>
                                </v-col>

                                <!-- </div> -->
                            </v-row>
                            <v-row dense v-for="(data,index) in ing" :key="index">
                                
                                <v-col class="ml-5" cols="3">
                                    {{data.name}}
                                </v-col>
                                <v-col  cols="3">
                                    {{data.source}}
                                </v-col>
                                <!-- <p>{{data.name}} : </p>
                                <p>{{data.source}}</p> -->
                            </v-row>
                            <v-row dense>
                                <v-col cols="2" class="checkbox">
                                    <v-checkbox label="Seafood?" color="primary" hide-details name="seafood" v-model="seafood">
                                    </v-checkbox>
                                </v-col>
                                <v-col cols="2">
                                    <v-checkbox label="Milk?" color="primary" hide-details name="milk" v-model="milk">
                                    </v-checkbox>
                                </v-col>
                                <v-col cols="2">
                                    <v-checkbox label="Peanut?" color="primary" hide-details name="peanut" v-model="peanut">
                                    </v-checkbox>
                                </v-col>
                                <v-col cols="2">
                                    <v-checkbox label="Egg?" color="primary" hide-details name="egg" v-model="egg">
                                    </v-checkbox>
                                </v-col>
                            </v-row>
                            <v-row dense>
                                <v-col cols="2" class="checkbox">
                                    <v-checkbox label="Is Promotion menu?" color="primary" hide-details name="promo" v-model="promo">
                                    </v-checkbox>
                                </v-col>
                                
                            </v-row>
                            <v-row dense>
                                <v-col cols="3" class="Select">
                                    <v-select :items="items" label="Choose type" outlined name="type" v-model="type"></v-select>
                                </v-col>
                            </v-row>
                            
  
                            <v-row dense>
                                <v-col col="2">
                                    <v-file-input placeholder="Upload your documents" label="Upload Pic" multiple
                                        prepend-icon="mdi-paperclip" class="Upload" @change="onImgSelected">
                                        <template v-slot:selection="{ text }">
                                            <v-chip small label color="primary">
                                                {{ text }}
                                            </v-chip>
                                        </template>
                                    </v-file-input>
                                </v-col>
                                <v-col col="2">
                                    <v-btn outlined rounded text color="green" dark @click="onUpload">
                                        Upload
                                    </v-btn>
                                </v-col>
                            </v-row>
                    
                            <v-row class="mt-5 mb-5 mr-5 d-flex flex-row-reverse">
                                <!-- <v-btn rounded color="grey darken-2" dark @click="onCreateIng">
                                    Test
                                </v-btn> -->
                                <v-btn :disabled="!checkup" rounded @click="onCreateIng()" class="mr-5" color="success">
                                    Create
                                </v-btn >
                                <v-btn rounded color="red" dark @click="onClear" class="mr-5">
                                    Clear
                                </v-btn>
                            </v-row>
                            <br>
                        </v-container>

                    </v-form>
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
        components:{
            StaffHeader,
            Header,
        },
        data: () => ({
            items: [{
                    text: 'Normal',
                    value: 1
                },
                {
                    text: 'Vegan',
                    value: 2
                },
                {
                    text: 'Vegetarian',
                    value: 3
                }
            ],
            return: {
                slider: 0,
                
            },
            select: null,
            checkup: false,
            menudish: null,
            price: 0,
            type: null,
            desc: null,
            status: true,
            promo:false,
            checkbox: false,
            seafood: false,
            milk: false,
            peanut: false,
            egg: false,
            ing:[],
            previous:[],
            ing_name:null,
            ing_source:null,
            count:1
            
        }),


        methods: {
            
            onImgSelected(event) {
                this.selectedFile = event[0]
                console.log(event[0])
            },
            
            onClear() {
                this.menudish = null
                this.price = null
                this.desc = null
                this.seafood = false
                this.milk = false
                this.peanut = false
                this.egg = false
                this.promo =false
                this.ing_name=null
                this.ing_source=null
                this.ing = []
            },

            onUpload() {
                this.checkup = true;
                const fd = new FormData();
                fd.append('image', this.selectedFile, this.selectedFile.name)
                axios.post(process.env.VUE_APP_BASE_URL+'/image/upload-image', fd)
                // axios.post('http://localhost:8081/image/upload-image', fd)
                    .then(res => {
                        console.log(res)
                    })
            },
            onUploadAds() {
                this.checkup = true;
                const fd = new FormData();
                fd.append('image', this.selectedFile, this.selectedFile.name)
                axios.post(process.env.VUE_APP_BASE_URL+'/pro/promotion-image', fd)
                // axios.post('http://localhost:8081/pro/promotion-image', fd)
                    .then(res => {
                        console.log(res)
                    })
            },

            async onCreate() {

                const fd = {
                    "name": this.menudish,
                    "price": this.price,
                    "description": this.desc,
                    "has_milk": this.milk,
                    "has_egg": this.egg,
                    "has_bean": this.peanut,
                    "has_seafood": this.seafood,
                    "type_id": this.type,
                    "spc_food":this.promo,
                    "status": this.status,
                };
                console.log(fd)
                const test = JSON.stringify(fd)
                console.log(test)
                

                axios.post(process.env.VUE_APP_BASE_URL+'/food/create-food', test)
                // axios.post('http://localhost:8081/food/create-food', test)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })
                this.checkup = false
                
            },

            onTest() {
                console.log(this.ing)
                console.log(this.ing_test)
                console.log(this.dialogm1)
            },
            onAddIng(name,src){
                this.ing.push({
                    name: name,
                    source : src
                })
                this.ing_name = ""
                this.ing_source = ""
                
            },
            
            async onCreateIng(){
                await this.onCreate();
                const test2 = JSON.stringify(this.ing)
                axios.post(process.env.VUE_APP_BASE_URL+'/food/create-ingredients', test2)
                // axios.post('http://localhost:8081/food/create-ingredients', test2)
                    .then(res => {
                        console.log(res)
                    }).catch(error => {
                        console.log(error)
                    })
                
            },
        },

    }
</script>

<style lang="scss">
    @import '@/styles/variables.scss';
    

    .checkbox {
        margin-left: 20px;
    }

    .Name {
        margin-left: 20px;
        margin-top: 20px;
    }

    .Select {
        margin-left: 20px;
        margin-top: 35px;
    }

    .Upload {
        margin-left: 20px
    }
    
    // .checkbox{
    //     checkbox-dense-margin-top	: 1px
    // }
</style>