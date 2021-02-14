<template>
  <!-- 一覧 -->
  <v-list two-line>
    <div v-for="item in items" :key="item.ID">
      <!-- アイテム -->
      <v-list-item>
        <!-- サイト画像 -->
        <v-list-item-avatar size="100">
          <v-img :src="item.ImageURL"></v-img>
        </v-list-item-avatar>
        <v-list-item-content>
          <!-- タイトル -->
          <v-list-item-title class="title" v-text="item.Title"></v-list-item-title>
          <!-- ディスクリプション -->
          <v-list-item-content>{{ item.Description }}</v-list-item-content>
          <div class="item-link-container">
            <a class="item-link" :href="item.URL" target="_blank">
              <v-icon>mdi-open-in-new</v-icon>
            </a>
          </div>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
    </div>
  </v-list>
</template>

<script>
export default {
  name: "SlideList",
  props: {
    items: {
      type: Array,
      default: () => ([]),
    },
  },
}
</script>

<style scoped>
.title {
  font-size: 1.3rem;
}

.item-link-container {
  display: flex;
  flex-direction: row-reverse;
}

.item-link {
  text-decoration: none
}
</style>