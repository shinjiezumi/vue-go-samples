<template>
  <!-- 一覧 -->
  <v-list two-line>
    <div v-for="item in items" :key="item.ID">
      <!-- アイテム -->
      <v-list-item>
        <v-list-item-content>
          <!-- ユーザー情報＋投稿日時 -->
          <div class="Item__UserInfo">
            <v-list-item-avatar class="item__UserAvatar" size="30">
              <v-img :src="item.ProfileImageURL" />
            </v-list-item-avatar>
            <div class="Item__CreatedMessage">
              <a href="'https://qiita.com/' + item.UserID" target="_blank">{{ item.UserID }}</a>
              <span> が {{ item.Created }} に投稿</span>
            </div>
          </div>
          <!-- タイトル -->
          <v-list-item-title class="Item__Title" v-text="item.Title" />
          <!-- タグ -->
          <div class="Item__Tags">
            <v-icon color="teal">
              mdi-tag
            </v-icon>
            <span v-for="(tag, i) in item.Tags" :key="tag" class="Item__Tag">
              {{ tag }}<template v-if="i !== (item.Tags.length - 1)">,</template>
            </span>
          </div>
          <!-- LGTM数 -->
          <div class="Item__Reaction">
            <v-icon color="teal">
              mdi-thumb-up
            </v-icon>
            <span class="Item__LikeCount">{{ item.LikeCount }}</span>
          </div>
          <!-- 記事リンク -->
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
  name: 'QiitaList',
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

.Item__UserInfo {
  display: flex;
  margin-top: -1rem;
}

.Item__CreatedMessage {
  margin-top: 20px;
}

.Item__Tags, .Item__Reaction {
  display: flex;
  margin-top: .5rem;
  word-break: keep-all;
}

.Item__Reaction {
  height: 1rem;
}

.Item__Tag, .Item__LikeCount {
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
