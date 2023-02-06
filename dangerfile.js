import { warn, message, markdown, danger } from "danger"

const { additions = 0, deletions = 0 } = danger.github.pr

const getModifiedFiles = () => danger.git.modified_files

const getDiffAdded = (fileName) => {
	danger.git.diffForFile(fileName).then((res) => {
		return res.added
	})
}

const ensurePRHasAssignee = () => {
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

const checkCommonIssue = () => {
	const modifiedFiles = getModifiedFiles()
	for (modifiedFile of modifiedFiles) {
		diffAddedForFile = getDiffAdded(modifiedFile) || ''
		if (diffAddedForFile.includes('fmt.Print')) {
			warn('fmt.print detected')
		}
	}
}
ensurePRHasAssignee()
reviewLargePR()
showCodeChanges()
checkCommonIssue()