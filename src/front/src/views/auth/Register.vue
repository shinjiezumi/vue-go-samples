<template>
  <v-main>
    <v-container>
      <v-row class="text-center" align="center" justify="center">
        <v-col cols="8">
          <v-alert v-if="this.getError !== ''" type="error" text>
            {{ this.getError }}
          </v-alert>
        </v-col>
      </v-row>
      <v-row class="text-center" align="center" justify="center">
        <v-col class="mb-4" cols="12" sm="8" md="4">
          <v-toolbar color="primary" dark>
            <v-toolbar-title>会員登録</v-toolbar-title>
          </v-toolbar>
          <v-card>
            <v-form>
              <v-col>
                <v-text-field
                  v-model="name"
                  type="text"
                  label="ユーザー名"
                  :error-messages="nameErrors"
                  required
                  @input="$v.name.$touch()"
                  @blur="$v.name.$touch()"
                />
              </v-col>
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
                <v-btn color="primary" @click="register">
                  会員登録
                </v-btn>
              </v-col>
              <v-col>
                アカウントをお持ちの方は
                <router-link to="/login">
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
import { validationMixin } from 'vuelidate';
import {
  email, maxLength, minLength, required,
} from 'vuelidate/lib/validators';
import { mapGetters } from 'vuex';
import { generateTitle } from '@/util';

export default {
  name: 'Login',
  title: generateTitle('会員登録'),
  mixins: [validationMixin],
  created() {
    this.$store.commit('error/clearError');
  },
  validations: {
    name: { required, maxLength: maxLength(255) },
    email: { required, email, maxLength: maxLength(255) },
    password: { required, minLength: minLength(8), maxLength: maxLength(16) },
  },
  data() {
    return {
      name: '',
      email: '',
      password: '',
    };
  },
  computed: {
    nameErrors() {
      const errors = [];
      if (!this.$v.name.$dirty) return errors;

      !this.$v.name.required && errors.push('ユーザー名を入力してください');
      !this.$v.name.maxLength && errors.push('ユーザー名は255文字以下で入力してください');

      return errors;
    },
    emailErrors() {
      const errors = [];
      if (!this.$v.email.$dirty) return errors;

      !this.$v.email.required && errors.push('メールアドレスを入力してください');
      !this.$v.email.email && errors.push('メールアドレスの形式が不正です');
      !this.$v.email.maxLength && errors.push('メールアドレスは255文字以下で入力してください');

      return errors;
    },
    passwordErrors() {
      const errors = [];
      if (!this.$v.password.$dirty) return errors;

      !this.$v.password.required && errors.push('パスワードを入力してください');
      (!this.$v.password.minLength || !this.$v.password.maxLength)
      && errors.push('パスワードは8文字以上16文字以下で入力してください');

      return errors;
    },
    ...mapGetters('error', ['getError']),
  },
  methods: {
    async register() {
      this.$v.$touch();
      if (this.$v.$invalid) return;

      await this.$store.dispatch('auth/register', {
        name: this.name,
        email: this.email,
        password: this.password,
      });

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
