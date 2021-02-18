<template>
  <!-- 一覧 -->
  <v-list two-line>
    <div v-for="item in items" :key="item.ID">
      <!-- アイテム -->
      <v-list-item>
        <v-list-item-content>
          <!-- タイトル -->
          <v-list-item-title class="title" v-text="item.Title"></v-list-item-title>
          <!-- PV+LGTM数} -->
          <div class="reaction-container">
            <v-icon color="teal">mdi-thumb-up</v-icon>
            <span class="like-count">{{ item.LikeCount }}</span>
          </div>
          <!-- タグ -->
          <div class="tag-container">
            <v-icon color="teal">mdi-tag</v-icon>
            <span v-for="(tag, i) in item.Tags" :key="tag" class="tag">
              {{ tag }}
              <span v-if="i !== (item.Tags.length - 1)">,</span>
            </span>
          </div>
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
  name: "QiitaList",
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

.tag-container, .reaction-container {
  display: flex;
  margin-top: .5rem;
}

.reaction-container {
  height: 1rem;
}

.tag, .like-count {
  padding: .2rem;
}

.item-link-container {
  display: flex;
  flex-direction: row-reverse;
}

.item-link {
  text-decoration: none
}
</style>