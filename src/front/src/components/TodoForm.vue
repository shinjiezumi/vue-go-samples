<template>
  <v-dialog v-model="dialog" persistent max-width="600px">
    <template v-slot:activator="{ on }">
      <v-row v-if="isNew" class="text-center" align="center" justify="center">
        <v-col class="mb-4" cols="12" sm="8">
          <v-btn class="mx-2" fab dark v-on="on" color="primary">
            <v-icon dark>mdi-plus</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-btn v-else class="mx-2" fab dark v-on="on" color="green" x-small>
        <v-icon dark>mdi-pencil</v-icon>
      </v-btn>
    </template>
    <v-card>
      <v-form>
        <v-col>
          <v-menu
              ref="menu1"
              v-model="menu1"
              :close-on-content-click="false"
              transition="scale-transition"
              offset-y
              max-width="290px"
              min-width="290px"
          >
            <template v-slot:activator="{ on }">
              <v-text-field v-model="dateFormatted" label="期限" v-on="on" @blur="date = parseDate(dateFormatted)"/>
            </template>
            <v-date-picker v-model="date" no-title @input="menu1 = false"/>
          </v-menu>
        </v-col>
        <v-col>
          <v-text-field
              type="text"
              v-model="title"
              label="タイトル"
              required
          />
        </v-col>
        <v-col>
          <v-textarea
              type="text"
              v-model="memo"
              label="メモ"
          />
        </v-col>
        <v-col>
          <v-card-actions>
            <v-spacer/>
            <v-btn color="blue darken-1" text @click="cancel">キャンセル</v-btn>
            <v-btn color="blue darken-1" text @click="store">保存</v-btn>
          </v-card-actions>
        </v-col>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script>
  import moment from "moment"
  import { formatDate, parseDate } from "../util";

  // TODO バリデーション実装
  export default {
    name: "TodoForm",
    props: {
      isNew: {
        type: Boolean,
        default: false,
      },
      todoId: {
        type: Number,
        default: 0,
      },
      todoTitle: {
        type: String,
        default: '',
      },
      todoMemo: {
        type: String,
        default: '',
      },
      todoLimitDate: {
        type: String,
        default: moment().format("YYYY-MM-DD"),
      },
    },
    data: vm => ({
      dialog: false,
      date: vm.todoLimitDate,
      dateFormatted: vm.formatDate(vm.todoLimitDate),
      title: vm.todoTitle,
      memo: vm.todoMemo,
      menu1: false,
    }),
    watch: {
      date () {
        this.dateFormatted = this.formatDate(this.date)
      },
    },
    computed: {
      getError() {
        return this.$store.getters['error/getError']
      }
    },
    methods: {
      formatDate(date) {
        return formatDate(date)
      },
      parseDate(date) {
        return parseDate(date)
      },
      getParams() {
        return {
          limit_date: this.date,
          title: this.title,
          memo: this.memo
        }
      },
      cancel() {
        this.clearForm(true)
      },
      async store() {
        if (!this.todoId) {
          await this.$store.dispatch('todo/create', {params: this.getParams()});
        } else {
          await this.$store.dispatch('todo/modify', {id: this.todoId, params: this.getParams()});
        }
        if (this.getError === '') {
          this.clearForm();
          this.$store.dispatch('todo/getList')
        }
      },
      clearForm(isCancel = false) {
        // TODO 編集時は都度fetchする
        if (!this.todoId) {
          this.title = "";
          this.memo = "";
          this.date = moment().format("YYYY-MM-DD")
        } else if (isCancel) {
          this.title = this.todoTitle;
          this.memo = this.todoMemo;
          this.date = this.todoLimitDate;
        }
        this.dialog = false;
      }
    },
  }
</script>

<style scoped>
</style>