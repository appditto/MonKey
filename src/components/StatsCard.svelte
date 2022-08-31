<script>
  export let header;
  export let stat;
  export let footer;
  function abbreviateNumber(value) {
    var newValue = value;
    if (value >= 1000) {
      var suffixes = ["", "K", "M", "B", "T"];
      var suffixNum = Math.floor(("" + value).length / 3);
      var shortValue = "";
      for (var precision = 3; precision >= 1; precision--) {
        shortValue = parseFloat(
          (suffixNum != 0 ? value / Math.pow(1000, suffixNum) : value).toPrecision(precision)
        );
        var dotLessShortValue = (shortValue + "").replace(/[^a-zA-Z 0-9]+/g, "");
        if (dotLessShortValue.length <= 3) {
          break;
        }
      }
      if (shortValue % 1 != 0)
        shortValue = shortValue.toLocaleString("en-US", {
          maximumFractionDigits: 2,
          minimumFractionDigits: 0,
        });
      newValue = shortValue + suffixes[suffixNum];
    }
    return newValue;
  }
</script>

<div class="w-full md:w-1/2 lg:w-1/3 max-w-sm px-2 py-3">
  <div class="bg-white flex flex-col justify-center items-center card pt-6 pb-8">
    <p>{header}</p>
    <p class="font-bold text-4xl">{abbreviateNumber(stat)}</p>
    <p>{footer}</p>
  </div>
</div>

<style>
  .card {
    border-radius: 0.65rem;
    border-width: 2px;
    border-color: #e4e5e9;
    box-shadow: 0rem 0.3rem 0rem 0rem #e4e5e9;
  }
</style>
