<script lang="ts">
	interface Props {
		div: string;
		styles?: string;
	}

	let { div, styles = '' }: Props = $props();
</script>

<span
	class="rounded-full px-2 py-1 text-2xl/4 border-division-{div} border-x-3 text-division-{div} bg-jfgray-900/75 {styles}"
	>{div}</span
>
