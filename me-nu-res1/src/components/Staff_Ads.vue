<template>
    <v-app>
        <Header />
        <v-main>
            <v-container>
                <div>
                    <StaffHeader />
                </div>
            </v-container>
            <v-card class="mx-auto mt-5" max-width="1000">
                <v-container fluid>
                    <v-row dense>
                        <v-col v-for="(ads,index) in ads" :key="index" :cols="4">
                            <v-card height="auto" width="auto">
                                <v-card-text>
                                    <v-img :src ="ads.image_pro" height="auto" width="auto">

                                    </v-img>
                                </v-card-text>

                                <v-card-actions>
                                    <v-spacer></v-spacer>

                                    <v-btn color="red" text @click="onDelete(menu.id)">
                                        <v-icon>mdi-trash-can</v-icon>

                                    </v-btn>
                                    <div class="text-center">
                                    </div>
                                </v-card-actions>
                            </v-card>
                        </v-col>
                    </v-row>
                </v-container>
            </v-card>
            
            <v-row  class="mt-5 ml-15">
                <v-col col="6">
                    <v-file-input placeholder="Upload your documents" label="Upload Promotion Pic" multiple
                        prepend-icon="mdi-paperclip" class="Upload" @change="onImgSelected">
                        <template v-slot:selection="{ text }">
                            <v-chip small label color="primary">
                                {{ text }}
                            </v-chip>
                        </template>
                    </v-file-input>
                </v-col>
                <v-col col="6">
                    <v-btn outlined rounded text color="green" dark @click="onUploadAds">
                        Upload Ads
                    </v-btn>
                </v-col>
            </v-row>
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
                ads: [],
                select: null,
            }
        },
        mounted() {

        },
        created() {
            this.getdata()
        },
        methods: {
            getdata() {
                axios
                    .get(process.env.VUE_APP_BASE_URL + '/food/cus/menu-food')
                    .then(res => {
                        this.ads = res.data.data.item.promotion
                        console.log(this.ads)
                    })
            },
            onUploadAds() {
                const fd = new FormData();
                fd.append('image', this.selectedFile, this.selectedFile.name)
                axios.post(process.env.VUE_APP_BASE_URL + '/pro/promotion-image', fd)
                    // axios.post('http://localhost:8081/pro/promotion-image', fd)
                    .then(res => {
                        console.log(res)
                    })
            },
            onImgSelected(event) {
                this.selectedFile = event[0]
                console.log(event[0])
            },
        }
    }
</script>


<style>

</style>