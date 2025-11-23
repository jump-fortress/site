<script lang="ts">
	let { route, parentRoute, pages }: { route: string; parentRoute: string; pages: string[] } =
		$props();
</script>

<div class="sticky top-16 z-20 bg-jfgray-800/90 backdrop-blur-sm">
	<div class="relative flex h-12 w-fit transition-colors hover:text-ctp-lavender-50/75">
		<hr
			class="absolute right-0 bottom-3 left-0 m-auto w-11/12 rounded-full border-1 border-jfgray-700"
		/>
		{#each pages as name}
			{@render FormatPage(name)}
		{/each}
	</div>
</div>

{#snippet FormatPage(name: string)}
	<a
		href="/{parentRoute}/{name}"
		class="flex h-full items-center px-2 transition-colors first:pl-0 hover:text-ctp-lavender-50"
		><div class="relative">
			<span>{name}</span>
			{#if route.includes(name)}
				<hr
					class="absolute right-0 left-0 m-auto flex w-11/12 rounded-full border-1 text-ctp-lavender"
				/>
			{/if}
		</div></a
	>
{/snippet}
