<script>
    import { json } from "svelte-i18n";
    import { onMount } from "svelte";
    import { pageTitle } from "@/stores/app";

    $pageTitle = $json("common.menu.oauth2");

    onMount(() => {
        window.close();
    });
</script>

<div class="content txt-hint txt-center p-base">
    <h3 class="m-b-sm">Auth failed.</h3>
    <h5>You can close this window and go back to the app to try again.</h5>
</div>
