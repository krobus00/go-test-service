import { warn, message, fail, danger } from "danger"

const { additions = 0, deletions = 0, title = "" } = danger.github.pr

const getModifiedFiles = () => danger.git.modified_files

const getDiffAdded = async (fileName) => {
	let diff = await danger.git.diffForFile(fileName)
	return diff.added
}

const ensurePRHasAssignee = () => {
	if (title.toLowerCase().includes("hotfix")) return
	if (danger.github.pr.assignee === null) {
		fail("Please assign someone to merge this PR, and optionally include people who should review.")
	}
}

const showCodeChanges = () => {
	message(`:tada: The PR added ${additions} and removed ${deletions} lines.`)
	if (additions > deletions) {
		message(`:thumbsup: You removed more code than added!`);
	}
}

const reviewLargePR = () => {
	const bigPRThreshold = 600;
	if (additions + deletions > bigPRThreshold) {
		warn(`:exclamation: Pull Request size seems relatively large. If Pull Request contains multiple changes, split each into separate PR for faster, easier review.`);
	}
}

const checkCommonIssue = async () => {
	const modifiedFiles = getModifiedFiles()
	for (modifiedFile of modifiedFiles) {
		diffAddedForFile = await getDiffAdded(modifiedFile) || ''
		if (diffAddedForFile.includes('fmt.Print')) {
			fail('fmt.Print detected')
		}
	}
}

(async () => {
	try {
		ensurePRHasAssignee()
		reviewLargePR()
		showCodeChanges()
		await checkCommonIssue()
	} catch (e) {
		console.log(e)
	}
})();
