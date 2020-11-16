package fun

import (
	"GithubEventNF/main/model"
	"fmt"
	"github.com/kataras/iris/v12"
)

func Hook(ctx iris.Context) {
	eventName := ctx.GetHeader("X-GitHub-Event")
	fmt.Println(eventName)
	switch eventName {
	case "check_run":
		var mod model.CheckRun
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		CheckRunFun(eventName, &mod)
	case "check_suite":
		var mod model.CheckSuite
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		CheckSuiteFun(eventName, &mod)
	case "code_scanning_alert":
		var mod model.CodeScanningAlert
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		CodeScanningAlertFun(eventName, &mod)
	case "commit_comment":
		var mod model.CommitComment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		CommitCommentFun(eventName, &mod)
	case "content_reference":
		var mod model.ContentReference
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		ContentReferenceFun(eventName, &mod)
	case "create":
		var mod model.Create
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		CreateFun(eventName, &mod)
	case "delete":
		var mod model.Delete
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		DeleteFun(eventName, &mod)
	case "deploy_key":
		var mod model.DeployKey
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		DeployKeyFun(eventName, &mod)
	case "deployment":
		var mod model.Deployment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		DeploymentFun(eventName, &mod)
	case "deployment_status":
		var mod model.DeploymentStatus
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		DeploymentStatusFun(eventName, &mod)
	case "fork":
		var mod model.Fork
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		ForkFun(eventName, &mod)
	case "github_app_authorization":
		var mod model.GithubAppAuthorization
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		GithubAppAuthorizationFun(eventName, &mod)
	case "gollum":
		var mod model.Gollum
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		GollumFun(eventName, &mod)
	case "installation":
		var mod model.Installation
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		InstallationFun(eventName, &mod)
	case "installation_repositories":
		var mod model.InstallationRepositories
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		InstallationRepositoriesFun(eventName, &mod)
	case "issue_comment":
		var mod model.IssueComment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		IssueCommentFun(eventName, &mod)
	case "issues":
		var mod model.Issues
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		IssuesFun(eventName, &mod)
	case "label":
		var mod model.Label
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		LabelFun(eventName, &mod)
	case "marketplace_purchase":
		var mod model.MarketplacePurchase
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		MarketplacePurchaseFun(eventName, &mod)
	case "member":
		var mod model.Member
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		MemberFun(eventName, &mod)
	case "membership":
		var mod model.Membership
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		MembershipFun(eventName, &mod)
	case "meta":
		var mod model.Meta
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		MetaFun(eventName, &mod)
	case "milestone":
		var mod model.Milestone
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		MilestoneFun(eventName, &mod)
	case "org_block":
		var mod model.OrgBlock
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		OrgBlockFun(eventName, &mod)
	case "organization":
		var mod model.Organization
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		OrganizationFun(eventName, &mod)
	case "package":
		var mod model.Package
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PackageFun(eventName, &mod)
	case "page_build":
		var mod model.PageBuild
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PageBuildFun(eventName, &mod)
	case "ping":
		var mod model.Ping
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PingFun(eventName, &mod)
	case "project":
		var mod model.Project
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		ProjectFun(eventName, &mod)
	case "project_card":
		var mod model.ProjectCard
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		ProjectCardFun(eventName, &mod)
	case "project_column":
		var mod model.ProjectColumn
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		ProjectColumnFun(eventName, &mod)
	case "public":
		var mod model.Public
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PublicFun(eventName, &mod)
	case "pull_request":
		var mod model.PullRequest
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PullRequestFun(eventName, &mod)
	case "pull_request_review":
		var mod model.PullRequestReview
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PullRequestReviewFun(eventName, &mod)
	case "pull_request_review_comment":
		var mod model.PullRequestReviewComment
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PullRequestReviewCommentFun(eventName, &mod)
	case "push":
		var mod model.Push
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		PushFun(eventName, &mod)
	case "release":
		var mod model.Release
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		ReleaseFun(eventName, &mod)
	case "repository":
		var mod model.Repository
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		RepositoryFun(eventName, &mod)
	case "repository_dispatch":
		var mod model.RepositoryDispatch
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		RepositoryDispatchFun(eventName, &mod)
	case "repository_import":
		var mod model.RepositoryImport
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		RepositoryImportFun(eventName, &mod)
	case "repository_vulnerability_alert":
		var mod model.RepositoryVulnerabilityAlert
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		RepositoryVulnerabilityAlertFun(eventName, &mod)
	case "security_advisory":
		var mod model.SecurityAdvisory
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		SecurityAdvisoryFun(eventName, &mod)
	case "sponsorship":
		var mod model.Sponsorship
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		SponsorshipFun(eventName, &mod)
	case "star":
		var mod model.Star
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		StarFun(eventName, &mod)
	case "status":
		var mod model.Status
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		StatusFun(eventName, &mod)
	case "team":
		var mod model.Team
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		TeamFun(eventName, &mod)
	case "team_add":
		var mod model.TeamAdd
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		TeamAddFun(eventName, &mod)
	case "watch":
		var mod model.Watch
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		WatchFun(eventName, &mod)
	case "workflow_dispatch":
		var mod model.WorkflowDispatch
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		WorkflowDispatchFun(eventName, &mod)
	case "workflow_run":
		var mod model.WorkflowRun
		err := ctx.ReadJSON(&mod)
		if err != nil {
			panic(err)
		}
		WorkflowRunFun(eventName, &mod)
	default:
	}
}
