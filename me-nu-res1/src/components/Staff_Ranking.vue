<template>
    <v-app>
        <Header />
        <v-main>
            <v-container>
                <div>
                    <StaffHeader />
                </div>
            </v-container>
            <v-card class="mx-auto mt-5" width="1000" elevation="2" outlined shaped>
                <v-card-title class="justify-center">Rank during 2 weeks</v-card-title>
                <v-divider></v-divider>
                <v-card-text v-for="(data,index) in ranking" :key="index" >
                    <p class="font-weight-black text-center">{{index+1}} : {{data.name}} Qty :  {{data.amount}}

                    <v-icon v-show="index==0" color="#FFD700" class="ml-3">fas fa-crown</v-icon>
                    <v-icon v-show="index==1" color="#C0C0C0" class="ml-3">fas fa-crown</v-icon>
                    <v-icon v-show="index==2" small color="#CD7F32" class="ml-3">fas fa-crown</v-icon>
                    </p>
                    <!-- {{index+1}} : {{data.name}} Qty :  {{data.amount}} -->
                </v-card-text>
            </v-card>
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
                ranking: []
            }
        },
        mounted() {
            axios
                .get(process.env.VUE_APP_BASE_URL+'/food/cus/menu-food')
                .then(res => {
                    this.ranking = res.data.data.item.ranking
                    console.log(this.ranking)
                    // console.log(this.menu)
                })
        },
    }
</script>


<style>

</style>