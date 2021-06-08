<template>
  <v-dialog v-model="dialog" persistent max-width="600px">
    <template v-slot:activator="{ on }">
      <v-row v-if="isNew" class="createBtn text-center" align="center" justify="center">
        <v-col class="mb-4" cols="12" sm="8">
          <v-btn class="mx-2" fab dark color="primary" v-on="on">
            <v-icon dark>
              mdi-plus
            </v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-btn
        v-else
        class="mx-2"
        fab
        dark
        color="green"
        x-small
        v-on="on"
      >
        <v-icon dark>
          mdi-pencil
        </v-icon>
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
              <v-text-field
                v-model="dateFormatted"
                label="期限"
                v-on="on"
                @blur="limitDate = parseDate(dateFormatted)"
              />
            </template>
            <v-date-picker v-model="limitDate" no-title @input="menu1 = false" />
          </v-menu>
        </v-col>
        <v-col>
          <v-text-field
            v-model="title"
            type="text"
            label="タイトル"
            required
            :error-messages="titleErrors"
            @input="$v.title.$touch()"
            @blur="$v.title.$touch()"
          />
        </v-col>
        <v-col>
          <v-textarea
            v-model="memo"
            type="text"
            label="メモ"
            :error-messages="memoErrors"
            @input="$v.memo.$touch()"
            @blur="$v.memo.$touch()"
          />
        </v-col>
        <v-col>
          <v-card-actions>
            <v-spacer />
            <v-btn color="blue darken-1" text @click="cancel">
              キャンセル
            </v-btn>
            <v-btn color="blue darken-1" text @click="store">
              保存
            </v-btn>
          </v-card-actions>
        </v-col>
      </v-form>
    </v-card>
  </v-dialog>
</template>

<script>
import { validationMixin } from 'vuelidate';
import moment from 'moment';
// eslint-disable-next-line import/no-unresolved
import { formatDate, parseDate } from '@/util';
import { maxLength, required } from 'vuelidate/lib/validators';

export default {
  name: 'TodoForm',
  mixins: [validationMixin],
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
      default: moment().format('YYYY-MM-DD'),
    },
  },
  validations: {
    title: { required, maxLength: maxLength(128) },
    memo: { maxLength: maxLength(255) },
  },
  data: (vm) => ({
    dialog: false,
    today: moment(),
    limitDate: vm.todoLimitDate,
    dateFormatted: vm.formatDate(vm.todoLimitDate),
    title: vm.todoTitle,
    memo: vm.todoMemo,
    menu1: false,
  }),
  computed: {
    error() {
      return this.$store.getters['error/getError'];
    },
    errorCode() {
      return this.$store.getters['error/getCode'];
    },
    titleErrors() {
      const errors = [];
      if (!this.$v.title.$dirty) return errors;

      // eslint-disable-next-line no-unused-expressions
      !this.$v.title.required && errors.push('タイトルを入力してください');
      // eslint-disable-next-line no-unused-expressions
      !this.$v.title.maxLength && errors.push('タイトルは128文字以下で入力してください');

      return errors;
    },
    memoErrors() {
      const errors = [];
      if (!this.$v.memo.$dirty) return errors;

      // eslint-disable-next-line no-unused-expressions
      !this.$v.memo.maxLength && errors.push('メモは255文字以下で入力してください');

      return errors;
    },
  },
  watch: {
    limitDate() {
      this.dateFormatted = this.formatDate(this.limitDate);
    },
  },
  methods: {
    formatDate(date) {
      return formatDate(date);
    },
    parseDate(date) {
      return parseDate(date);
    },
    getParams() {
      return {
        limit_date: this.limitDate,
        title: this.title,
        memo: this.memo,
      };
    },
    cancel() {
      this.clearForm(true);
    },
    async store() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      if (!this.todoId) {
        await this.$store.dispatch('todo/create', { params: this.getParams() });
      } else {
        await this.$store.dispatch('todo/modify', { id: this.todoId, params: this.getParams() });
      }
      if (this.error === '') {
        this.$emit('get-todo-list');
      } else {
        this.$emit('handle-error');
      }

      this.clearForm();
    },
    clearForm(isCancel = false) {
      // TODO 編集時は都度fetchする
      if (!this.todoId) {
        this.title = '';
        this.memo = '';
        this.limitDate = moment().format('YYYY-MM-DD');
      } else if (isCancel) {
        this.title = this.todoTitle;
        this.memo = this.todoMemo;
        this.limitDate = this.todoLimitDate;
      }
      this.$v.$reset();
      this.dialog = false;
    },
  },
};
</script>

<style scoped>
.createBtn {
  margin-top: .5rem;
}
</style>
