<script>
  import StatsCard from "../StatsCard.svelte";
  export let stats;
  export let statsMonthly;
</script>

<!-- HTML -->
<div class="w-full pt-16 pb-24 md:pb-40">
  <div class="y-container flex flex-col items-center">
    <h2 class="font-bold text-4xl md:text-5xl mx-4 text-center leading-tight tracking-tight mt-5">
      MonKey Stats
    </h2>
    <div class="w-full flex flex-row flex-wrap justify-center mt-5 md:mt-8 px-6 md:px-12">
      <StatsCard header="We served" stat={stats.total_served} footer="MonKeys in total" />
      <StatsCard
        header="We served"
        stat={statsMonthly.last30address.total}
        footer="MonKeys this month"
      />
      <StatsCard
        header="We served to"
        stat={stats.unique_clients_served}
        footer="clients in total"
      />
      <StatsCard
        header="We served to"
        stat={statsMonthly.last30clients.unique}
        footer="clients this month"
      />
      <StatsCard
        header="Available on"
        stat={Object.keys(stats.services).length - 1}
        footer="services"
      />
    </div>
  </div>
</div>
