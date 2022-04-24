<template>
  <v-card class="mx-auto my-5" max-width="800" max-height="600">
     <v-carousel cycle height="200" hide-delimiter-background :show-arrows="false">
      <v-carousel-item v-for="(ads, i) in ads" :key="i" :src="ads.image_pro">
        <!-- <v-sheet :color="colors[i]" height="100%">
          <v-row class="fill-height" align="center" justify="center">
            <div class="text-h2">
              {{ slide }}
            </div>
          </v-row>
        </v-sheet> -->

      </v-carousel-item>
    </v-carousel>
  </v-card>
</template>

<script>
  export default {
    props:["ads"],
    data() {
      return {
        colors: [
          'indigo',
          'warning',
          'pink darken-2',
          'red lighten-1',
          'deep-purple accent-4',
        ],
        slides: [
          'เมนูอาหารประจำเดือนรายการที่ 1',
          'เมนูอาหารประจำเดือนรายการที่ 2',
          'เมนูอาหารประจำเดือนรายการที่ 3',
          'เมนูอาหารประจำเดือนรายการที่ 4',
          'เมนูอาหารประจำเดือนรายการที่ 5',
          
        ],
      }
    },
  }
</script>

<style>

</style>