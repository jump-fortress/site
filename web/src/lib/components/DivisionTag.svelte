<script lang="ts">
	interface Props {
		div: string;
		style?: string;
	}

	let { div, style = '' }: Props = $props();
</script>

<span
	class="rounded-full px-2 pb-0.5 pt-1 text-2xl/4 border-division-{div} border-x-3 text-division-{div} bg-jfgray-900/75 {style}"
	>{div}</span
>
