package fun

import (
	"GithubEventNF/main/model"
	"github.com/kataras/iris/v12"
	"log"
)

func Hook(ctx iris.Context) {
	eventName := ctx.GetHeader("X-GitHub-Event")
	log.Println(eventName)
	switch eventName {
	case "check_run":
		var mod model.CheckRun
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		CheckRunFun(eventName, &mod)
	case "check_suite":
		var mod model.CheckSuite
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		CheckSuiteFun(eventName, &mod)
	case "code_scanning_alert":
		var mod model.CodeScanningAlert
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		CodeScanningAlertFun(eventName, &mod)
	case "commit_comment":
		var mod model.CommitComment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		CommitCommentFun(eventName, &mod)
	case "content_reference":
		var mod model.ContentReference
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		ContentReferenceFun(eventName, &mod)
	case "create":
		var mod model.Create
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		CreateFun(eventName, &mod)
	case "delete":
		var mod model.Delete
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		DeleteFun(eventName, &mod)
	case "deploy_key":
		var mod model.DeployKey
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		DeployKeyFun(eventName, &mod)
	case "deployment":
		var mod model.Deployment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		DeploymentFun(eventName, &mod)
	case "deployment_status":
		var mod model.DeploymentStatus
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		DeploymentStatusFun(eventName, &mod)
	case "fork":
		var mod model.Fork
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		ForkFun(eventName, &mod)
	case "github_app_authorization":
		var mod model.GithubAppAuthorization
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		GithubAppAuthorizationFun(eventName, &mod)
	case "gollum":
		var mod model.Gollum
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		GollumFun(eventName, &mod)
	case "installation":
		var mod model.Installation
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		InstallationFun(eventName, &mod)
	case "installation_repositories":
		var mod model.InstallationRepositories
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		InstallationRepositoriesFun(eventName, &mod)
	case "issue_comment":
		var mod model.IssueComment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		IssueCommentFun(eventName, &mod)
	case "issues":
		var mod model.Issues
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		IssuesFun(eventName, &mod)
	case "label":
		var mod model.Label
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		LabelFun(eventName, &mod)
	case "marketplace_purchase":
		var mod model.MarketplacePurchase
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		MarketplacePurchaseFun(eventName, &mod)
	case "member":
		var mod model.Member
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		MemberFun(eventName, &mod)
	case "membership":
		var mod model.Membership
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		MembershipFun(eventName, &mod)
	case "meta":
		var mod model.Meta
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		MetaFun(eventName, &mod)
	case "milestone":
		var mod model.Milestone
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		MilestoneFun(eventName, &mod)
	case "org_block":
		var mod model.OrgBlock
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		OrgBlockFun(eventName, &mod)
	case "organization":
		var mod model.Organization
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		OrganizationFun(eventName, &mod)
	case "package":
		var mod model.Package
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PackageFun(eventName, &mod)
	case "page_build":
		var mod model.PageBuild
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PageBuildFun(eventName, &mod)
	case "ping":
		var mod model.Ping
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PingFun(eventName, &mod)
	case "project":
		var mod model.Project
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		ProjectFun(eventName, &mod)
	case "project_card":
		var mod model.ProjectCard
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		ProjectCardFun(eventName, &mod)
	case "project_column":
		var mod model.ProjectColumn
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		ProjectColumnFun(eventName, &mod)
	case "public":
		var mod model.Public
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PublicFun(eventName, &mod)
	case "pull_request":
		var mod model.PullRequest
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PullRequestFun(eventName, &mod)
	case "pull_request_review":
		var mod model.PullRequestReview
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PullRequestReviewFun(eventName, &mod)
	case "pull_request_review_comment":
		var mod model.PullRequestReviewComment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PullRequestReviewCommentFun(eventName, &mod)
	case "push":
		var mod model.Push
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		PushFun(eventName, &mod)
	case "release":
		var mod model.Release
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		ReleaseFun(eventName, &mod)
	case "repository":
		var mod model.Repository
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		RepositoryFun(eventName, &mod)
	case "repository_dispatch":
		var mod model.RepositoryDispatch
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		RepositoryDispatchFun(eventName, &mod)
	case "repository_import":
		var mod model.RepositoryImport
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		RepositoryImportFun(eventName, &mod)
	case "repository_vulnerability_alert":
		var mod model.RepositoryVulnerabilityAlert
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		RepositoryVulnerabilityAlertFun(eventName, &mod)
	case "security_advisory":
		var mod model.SecurityAdvisory
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		SecurityAdvisoryFun(eventName, &mod)
	case "sponsorship":
		var mod model.Sponsorship
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		SponsorshipFun(eventName, &mod)
	case "star":
		var mod model.Star
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		StarFun(eventName, &mod)
	case "status":
		var mod model.Status
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		StatusFun(eventName, &mod)
	case "team":
		var mod model.Team
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		TeamFun(eventName, &mod)
	case "team_add":
		var mod model.TeamAdd
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		TeamAddFun(eventName, &mod)
	case "watch":
		var mod model.Watch
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		WatchFun(eventName, &mod)
	case "workflow_dispatch":
		var mod model.WorkflowDispatch
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		WorkflowDispatchFun(eventName, &mod)
	case "workflow_run":
		var mod model.WorkflowRun
		err := ctx.ReadJSON(&mod)
		if err != nil {
			log.Panic(err)
		}
		WorkflowRunFun(eventName, &mod)
	default:
	}
}
