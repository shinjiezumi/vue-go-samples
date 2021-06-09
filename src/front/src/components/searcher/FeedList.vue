<template>
  <!-- 一覧 -->
  <v-list two-line>
    <div v-for="item in items" :key="item.ID">
      <!-- アイテム -->
      <v-list-item>
        <!-- サイト画像 -->
        <v-list-item-avatar size="100">
          <v-img :src="item.ImageURL" />
        </v-list-item-avatar>
        <v-list-item-content>
          <!-- タイトル -->
          <v-list-item-title class="Item__Title" v-text="item.Title" />
          <!-- タグ -->
          <div class="Item__Tags">
            <v-icon color="teal">
              mdi-tag
            </v-icon>
            <span v-for="(tag, i) in item.Tags" :key="tag" class="Item__Tag">
              {{ tag }}
              <span v-if="i !== (item.Tags.length - 1)">,</span>
            </span>
          </div>
          <!-- ディスクリプション -->
          <v-list-item-content>{{ item.Description }}</v-list-item-content>
          <div class="Item__Link">
            <a class="Item__LinkUrl" :href="item.URL" target="_blank">
              <v-icon>mdi-open-in-new</v-icon>
            </a>
          </div>
        </v-list-item-content>
      </v-list-item>
      <v-divider />
    </div>
  </v-list>
</template>

<script>
export default {
  name: 'FeedList',
  props: {
    items: {
      type: Array,
      default: () => ([]),
    },
  },
};
</script>

<style scoped>
.Item__Title {
  font-size: 1.3rem;
}

.Item__Tags {
  display: flex;
  margin-top: .5rem;
}

.Item__Tag {
  padding: .2rem;
}

.Item__Link {
  display: flex;
  flex-direction: row-reverse;
}

.Item__LinkUrl {
  text-decoration: none
}
</style>
