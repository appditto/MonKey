<script>
  import Hero from "../components/Hero.svelte";
  import How from "../components/How.svelte";
  import Rigged from "../components/Rigged.svelte";
</script>

<style global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  .y-container {
    max-width: 1280px;
    margin-left: auto;
    margin-right: auto;
  }
</style>

<svelte:head />

<!-- HTML -->
<Hero />
<How />
<Rigged />
