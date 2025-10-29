<script>
	import CompetitionPreview from '$lib/components/CompetitionPreview.svelte';
	import DataSection from '$lib/components/DataSection.svelte';
</script>

<DataSection title="upcoming / active monthly">
	<CompetitionPreview />
</DataSection>

<DataSection title="past monthlies">
	<CompetitionPreview small={true} />
</DataSection>
