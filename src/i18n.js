import { init, getLocaleFromNavigator, addMessages } from 'svelte-i18n';

import en from './locales/en.json'
import zh from './locales/zh.json'
import ja from './locales/ja.json'

addMessages("en", en);
addMessages("ja", ja);
addMessages("zh", zh);

init({
    fallbackLocale: 'en',
    initialLocale: getLocaleFromNavigator()
});