<template>
  <v-main>
    <v-container>
      <v-row class="text-center" align="center" justify="center">
        <v-col cols="8">
          <v-alert v-if="getError !== ''" type="error" text>
            {{ getError }}
          </v-alert>
        </v-col>
      </v-row>
      <v-row class="text-center" align="center" justify="center">
        <v-col class="mb-4" cols="12" sm="8" md="4">
          <v-toolbar color="primary" dark>
            <v-toolbar-title>ログイン</v-toolbar-title>
          </v-toolbar>
          <v-card>
            <v-form>
              <v-col>
                <v-text-field
                  v-model="email"
                  type="email"
                  label="メールアドレス"
                  :error-messages="emailErrors"
                  required
                  @input="$v.email.$touch()"
                  @blur="$v.email.$touch()"
                />
              </v-col>
              <v-col>
                <v-text-field
                  v-model="password"
                  type="password"
                  label="パスワード"
                  :error-messages="passwordErrors"
                  required
                  @input="$v.password.$touch()"
                  @blur="$v.password.$touch()"
                />
              </v-col>
              <v-col>
                <v-btn color="primary" @click="login">
                  ログイン
                </v-btn>
              </v-col>
              <v-col>
                <v-btn color="primary" @click="testLogin">
                  テストユーザーでログイン
                </v-btn>
              </v-col>
              <v-col>
                会員登録は
                <router-link to="/register">
                  こちら
                </router-link>
                から
              </v-col>
            </v-form>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import { mapGetters } from 'vuex';
import { validationMixin } from 'vuelidate';
import { email, minLength, required } from 'vuelidate/lib/validators';
import { generateTitle } from '@/util';

export default {
  name: 'Login',
  title: generateTitle('ログイン'),
  mixins: [validationMixin],
  created() {
    this.$store.commit('error/clearError');
  },
  validations: {
    email: { required, email },
    password: { required, minLength: minLength(8) },
  },
  computed: {
    emailErrors() {
      const errors = [];
      if (!this.$v.email.$dirty) return errors;
      !this.$v.email.required && errors.push('メールアドレスを入力してください');
      !this.$v.email.email && errors.push('メールアドレスの形式が不正です');
      return errors;
    },
    passwordErrors() {
      const errors = [];
      if (!this.$v.password.$dirty) return errors;
      !this.$v.password.required && errors.push('パスワードを入力してください');
      !this.$v.password.minLength && errors.push('パスワードは8文字以上で入力してください');
      return errors;
    },
    ...mapGetters('error', ['getError']),
  },
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    async login() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      await this.$store.dispatch('auth/login', { email: this.email, password: this.password });

      if (this.getError === '') {
        const getUser = async () => {
          await this.$store.dispatch('auth/currentUser');
          this.$router.push('/todos');
        };

        getUser();
      }
    },
    async testLogin() {
      await this.$store.dispatch('auth/login', { email: 'test@shinjiezumi.work', password: 'testtest' });

      if (this.getError === '') {
        const getUser = async () => {
          await this.$store.dispatch('auth/currentUser');
          this.$router.push('/todos');
        };

        getUser();
      }
    },
  },
};
</script>

<style scoped>

</style>
