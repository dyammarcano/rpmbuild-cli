<template>
    <div>
        <h2>Insurance Eligibility Form</h2>
        <form @submit.prevent="submitForm">
            <label>
                <input type="checkbox" v-model="healthIssues.smoke" /> Do you smoke?
            </label>
            <label>
                <input type="checkbox" v-model="healthIssues.drink" /> Do you drink alcohol?
            </label>
            <label>
                <input type="checkbox" v-model="healthIssues.noExercise" /> Do you not exercise regularly?
            </label>
            <!-- Add more checkboxes for other health issues if needed -->

            <button type="submit">Check Eligibility</button>
        </form>
        <div v-if="isEligible !== null">
            <p v-if="isEligible">Congratulations, you are eligible for insurance!</p>
            <p v-else>Sorry, you are not eligible for insurance based on your health issues.</p>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            healthIssues: {
                smoke: false,
                drink: false,
                noExercise: false,
                // Initialize other health issue properties here
            },
            isEligible: null,
        };
    },
    methods: {
        submitForm() {
            // Perform eligibility check based on health issues
            // For example, if any of the checkboxes are selected, mark as ineligible
            const ineligibleIssues = Object.values(this.healthIssues).some(issue => issue);
            this.isEligible = !ineligibleIssues;
        },
    },
};
</script>

<style scoped>
/* Add your component styles here */
</style>
