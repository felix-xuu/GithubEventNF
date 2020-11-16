package fun

import (
	"GithubEventNF/main/model"
	"strconv"
	"strings"
)

func CheckRunFun(eventName string, mod *model.CheckRun) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.CheckRun.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func CheckSuiteFun(eventName string, mod *model.CheckSuite) {
	var detail strings.Builder
	detail.WriteString("Branch => ")
	detail.WriteString(mod.CheckSuite.HeadBranch)
	detail.WriteString("\nStatus => ")
	detail.WriteString(mod.CheckSuite.Status)
	detail.WriteString("\nConclusion => ")
	detail.WriteString(mod.CheckSuite.Conclusion)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func CodeScanningAlertFun(eventName string, mod *model.CodeScanningAlert) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Alert.HtmlUrl)
	var user, userUrl string
	if "" == mod.Sender.Login {
		user = mod.Repository.Owner.Login
		userUrl = mod.Repository.Owner.HtmlUrl
	} else {
		user = mod.Sender.Login
		userUrl = mod.Sender.HtmlUrl
	}
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    user,
		userUrl: userUrl,
		detail:  detail.String(),
	})
}

func CommitCommentFun(eventName string, mod *model.CommitComment) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Comment.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func ContentReferenceFun(eventName string, mod *model.ContentReference) {
	var detail strings.Builder
	detail.WriteString("Reference => ")
	detail.WriteString(mod.ContentReference.Reference)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func CreateFun(eventName string, mod *model.Create) {
	var detail strings.Builder
	detail.WriteString("Ref => ")
	detail.WriteString(mod.Ref)
	detail.WriteString("\nRef Type => ")
	detail.WriteString(mod.RefType)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  "created",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func DeleteFun(eventName string, mod *model.Delete) {
	var detail strings.Builder
	detail.WriteString("Ref => ")
	detail.WriteString(mod.Ref)
	detail.WriteString("\nRef Type => ")
	detail.WriteString(mod.RefType)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  "deleted",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func DeployKeyFun(eventName string, mod *model.DeployKey) {
	var detail strings.Builder
	detail.WriteString("Key Title => ")
	detail.WriteString(mod.Key.Title)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func DeploymentFun(eventName string, mod *model.Deployment) {
	var detail strings.Builder
	detail.WriteString("Ref => ")
	detail.WriteString(mod.Deployment.Ref)
	detail.WriteString("\nEnvironment => ")
	detail.WriteString(mod.Deployment.Environment)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.Deployment.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func DeploymentStatusFun(eventName string, mod *model.DeploymentStatus) {
	var detail strings.Builder
	detail.WriteString("State => ")
	detail.WriteString(mod.DeploymentStatus.State)
	detail.WriteString("\nEnvironment => ")
	detail.WriteString(mod.DeploymentStatus.Environment)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.DeploymentStatus.Description)
	detail.WriteString("\nTarget Url => ")
	detail.WriteString(mod.DeploymentStatus.TargetUrl)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func ForkFun(eventName string, mod *model.Fork) {
	var detail strings.Builder
	detail.WriteString("Repo => ")
	detail.WriteString(mod.Forkee.FullName)
	detail.WriteString("\nFork Url => ")
	detail.WriteString(mod.Forkee.HtmlUrl)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  "fork",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func GithubAppAuthorizationFun(eventName string, mod *model.GithubAppAuthorization) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Sender.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func GollumFun(eventName string, mod *model.Gollum) {
	var detail strings.Builder
	for _, item := range mod.Pages {
		detail.WriteString("\nUrl => ")
		detail.WriteString(item.HtmlUrl)
	}
	SendMessage(Message{
		event:   eventName,
		action:  "Wiki page updated",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func InstallationFun(eventName string, mod *model.Installation) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Installation.HtmlUrl)
	var repos strings.Builder
	for _, item := range mod.Repositories {
		repos.WriteString("[")
		repos.WriteString(item.FullName)
		repos.WriteString("](")
		repos.WriteString(item.HtmlUrl)
		repos.WriteString("); ")
	}
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    repos.String(),
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func InstallationRepositoriesFun(eventName string, mod *model.InstallationRepositories) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Installation.HtmlUrl)
	var operateRepos []model.RepositoryDefine
	if "added" == mod.Action {
		operateRepos = mod.RepositoriesAdded
	} else {
		operateRepos = mod.RepositoriesRemoved
	}
	var repos strings.Builder
	for _, item := range operateRepos {
		repos.WriteString("[")
		repos.WriteString(item.FullName)
		repos.WriteString("](")
		repos.WriteString(item.HtmlUrl)
		repos.WriteString("); ")
	}
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    repos.String(),
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func IssueCommentFun(eventName string, mod *model.IssueComment) {
	var detail strings.Builder
	detail.WriteString("Issue => ")
	detail.WriteString(mod.Issue.Title)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Comment.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func IssuesFun(eventName string, mod *model.Issues) {
	var detail strings.Builder
	detail.WriteString("Issue => ")
	detail.WriteString(mod.Issue.Title)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Issue.HtmlUrl)
	switch mod.Action {
	case "assigned", "unassigned":
		detail.WriteString("\nAssignees => ")
		var assignees []string
		for _, item := range mod.Issue.Assignees {
			assignees = append(assignees, item.Login)
		}
		detail.WriteString(strings.Join(assignees, "; "))
	case "labeled", "unlabeled":
		detail.WriteString("\nLabels => ")
		var labels []string
		for _, item := range mod.Issue.Labels {
			labels = append(labels, item.Name)
		}
		detail.WriteString(strings.Join(labels, "; "))
	case "milestoned":
		detail.WriteString("\nMilestone Title => ")
		detail.WriteString(mod.Issue.Milestone.Title)
		detail.WriteString("\nMilestone Description => ")
		detail.WriteString(mod.Issue.Milestone.Description)
		detail.WriteString("\nMilestone Url => ")
		detail.WriteString(mod.Issue.Milestone.HtmlUrl)
	default:
	}
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func LabelFun(eventName string, mod *model.Label) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func MarketplacePurchaseFun(eventName string, mod *model.MarketplacePurchase) {
	var detail strings.Builder
	detail.WriteString("Billing Cycle => ")
	detail.WriteString(mod.MarketplacePurchase.BillingCycle)
	detail.WriteString("\nUnit Count => ")
	detail.WriteString(string(mod.MarketplacePurchase.UnitCount))
	detail.WriteString("\nOn Free Trial => ")
	detail.WriteString(strconv.FormatBool(mod.MarketplacePurchase.OnFreeTrial))
	detail.WriteString("\nNext Billing Date => ")
	detail.WriteString(mod.MarketplacePurchase.NextBillingDate)
	detail.WriteString("\nPlan Name => ")
	detail.WriteString(mod.MarketplacePurchase.Plan.Name)
	detail.WriteString("\nPlan Description => ")
	detail.WriteString(mod.MarketplacePurchase.Plan.Description)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func MemberFun(eventName string, mod *model.Member) {
	var detail strings.Builder
	detail.WriteString("Member Name => ")
	detail.WriteString(mod.Member.Login)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Member.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func MembershipFun(eventName string, mod *model.Membership) {
	var detail strings.Builder
	detail.WriteString("Scope => ")
	detail.WriteString(mod.Scope)
	detail.WriteString("\nMember Name => ")
	detail.WriteString(mod.Member.Login)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Member.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func MetaFun(eventName string, mod *model.Meta) {
	var detail strings.Builder
	detail.WriteString("Hook Name => ")
	detail.WriteString(mod.Hook.Name)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Hook.Config.Url)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func MilestoneFun(eventName string, mod *model.Milestone) {
	var detail strings.Builder
	detail.WriteString("Title => ")
	detail.WriteString(mod.Milestone.Title)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.Milestone.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Milestone.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func OrganizationFun(eventName string, mod *model.Organization) {
	var detail strings.Builder
	detail.WriteString("Aim User => ")
	detail.WriteString(mod.Membership.User.Login)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Membership.User.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func OrgBlockFun(eventName string, mod *model.OrgBlock) {
	var detail strings.Builder
	detail.WriteString("Aim User => ")
	detail.WriteString(mod.BlockedUser.Login)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.BlockedUser.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PackageFun(eventName string, mod *model.Package) {
	var detail strings.Builder
	detail.WriteString("Package Name => ")
	detail.WriteString(mod.Package.Name)
	detail.WriteString("\nPackage Version => ")
	detail.WriteString(mod.Package.PackageVersion.Version)
	detail.WriteString("\nPackage Summary => ")
	detail.WriteString(mod.Package.PackageVersion.Summary)
	detail.WriteString("\nPackage Body => ")
	detail.WriteString(mod.Package.PackageVersion.Body)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Package.PackageVersion.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PageBuildFun(eventName string, mod *model.PageBuild) {
	var detail strings.Builder
	detail.WriteString("Status => ")
	detail.WriteString(mod.Build.Status)
	detail.WriteString("\nError Message => ")
	detail.WriteString(mod.Build.Error.Message)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PingFun(eventName string, mod *model.Ping) {
	var detail strings.Builder
	detail.WriteString("Hook Name => ")
	detail.WriteString(mod.Hook.Name)
	detail.WriteString("\nEvents => ")
	detail.WriteString(strings.Join(mod.Hook.Events, "; "))
	SendMessage(Message{
		event:   eventName,
		action:  "webhook created",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func ProjectCardFun(eventName string, mod *model.ProjectCard) {
	var detail strings.Builder
	detail.WriteString("Note => ")
	detail.WriteString(mod.ProjectCard.Note)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func ProjectColumnFun(eventName string, mod *model.ProjectColumn) {
	var detail strings.Builder
	detail.WriteString("Name => ")
	detail.WriteString(mod.ProjectColumn.Name)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func ProjectFun(eventName string, mod *model.Project) {
	var detail strings.Builder
	detail.WriteString("Name => ")
	detail.WriteString(mod.Project.Name)
	detail.WriteString("\nBody => ")
	detail.WriteString(mod.Project.Body)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Project.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PublicFun(eventName string, mod *model.Public) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  "public repository",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PullRequestFun(eventName string, mod *model.PullRequest) {
	var detail strings.Builder
	detail.WriteString("State => ")
	detail.WriteString(mod.PullRequest.State)
	detail.WriteString("\nTitle => ")
	detail.WriteString(mod.PullRequest.Title)
	detail.WriteString("\nBody => ")
	detail.WriteString(mod.PullRequest.Body)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.PullRequest.HtmlUrl)
	detail.WriteString("\nDiff Url => ")
	detail.WriteString(mod.PullRequest.DiffUrl)
	detail.WriteString("\nPatch Url => ")
	detail.WriteString(mod.PullRequest.PatchUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PullRequestReviewFun(eventName string, mod *model.PullRequestReview) {
	var detail strings.Builder
	detail.WriteString("Pull Request Title => ")
	detail.WriteString(mod.PullRequest.Title)
	detail.WriteString("\nState => ")
	detail.WriteString(mod.Review.State)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Review.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PullRequestReviewCommentFun(eventName string, mod *model.PullRequestReviewComment) {
	var detail strings.Builder
	detail.WriteString("Pull Request Title => ")
	detail.WriteString(mod.PullRequest.Title)
	detail.WriteString("\nPath => ")
	detail.WriteString(mod.Comment.Path)
	detail.WriteString("\nBody => ")
	detail.WriteString(mod.Comment.Body)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Comment.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func PushFun(eventName string, mod *model.Push) {
	var detail strings.Builder
	detail.WriteString("Commits => ")
	for _, item := range mod.Commits {
		detail.WriteString("\n[")
		detail.WriteString(item.Id)
		detail.WriteString("](")
		detail.WriteString(item.Url)
		detail.WriteString(")  ")
		detail.WriteString(item.Message)
	}
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Compare)
	SendMessage(Message{
		event:   eventName,
		action:  "code pushed",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func ReleaseFun(eventName string, mod *model.Release) {
	var detail strings.Builder
	detail.WriteString("Name => ")
	detail.WriteString(mod.Release.Name)
	detail.WriteString("\nTag Name => ")
	detail.WriteString(mod.Release.TagName)
	detail.WriteString("\nBody => ")
	detail.WriteString(mod.Release.Body)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Release.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func RepositoryDispatchFun(eventName string, mod *model.RepositoryDispatch) {
	var detail strings.Builder
	detail.WriteString("Branch => ")
	detail.WriteString(mod.Branch)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func RepositoryFun(eventName string, mod *model.Repository) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func RepositoryImportFun(eventName string, mod *model.RepositoryImport) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Status,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func RepositoryVulnerabilityAlertFun(eventName string, mod *model.RepositoryVulnerabilityAlert) {
	var detail strings.Builder
	detail.WriteString("Affected Range => ")
	detail.WriteString(mod.Alert.AffectedRange)
	detail.WriteString("\nAffected Package Name => ")
	detail.WriteString(mod.Alert.AffectedPackageName)
	detail.WriteString("\nExternal Reference => ")
	detail.WriteString(mod.Alert.ExternalReference)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func SecurityAdvisoryFun(eventName string, mod *model.SecurityAdvisory) {
	var detail strings.Builder
	detail.WriteString("Summary => ")
	detail.WriteString(mod.SecurityAdvisory.Summary)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.SecurityAdvisory.Description)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func SponsorshipFun(eventName string, mod *model.Sponsorship) {
	var detail strings.Builder
	detail.WriteString("Tier Name => ")
	detail.WriteString(mod.Sponsorship.Tier.Name)
	detail.WriteString("\nTier Description => ")
	detail.WriteString(mod.Sponsorship.Tier.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Sponsorship.Sponsorable.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func StarFun(eventName string, mod *model.Star) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func StatusFun(eventName string, mod *model.Status) {
	var detail strings.Builder
	detail.WriteString("Description => ")
	detail.WriteString(mod.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  "commit changed to " + mod.State,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func TeamFun(eventName string, mod *model.Team) {
	var detail strings.Builder
	detail.WriteString("Name => ")
	detail.WriteString(mod.Team.Name)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.Team.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Team.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func TeamAddFun(eventName string, mod *model.TeamAdd) {
	var detail strings.Builder
	detail.WriteString("Name => ")
	detail.WriteString(mod.Team.Name)
	detail.WriteString("\nDescription => ")
	detail.WriteString(mod.Team.Description)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Team.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  "repository added to team",
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func WatchFun(eventName string, mod *model.Watch) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func WorkflowDispatchFun(eventName string, mod *model.WorkflowDispatch) {
	var detail strings.Builder
	detail.WriteString("Name => ")
	detail.WriteString(mod.Inputs.Name)
	detail.WriteString("\nWorkflow => ")
	detail.WriteString(mod.Workflow)
	detail.WriteString("\nUrl => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}

func WorkflowRunFun(eventName string, mod *model.WorkflowRun) {
	var detail strings.Builder
	detail.WriteString("Url => ")
	detail.WriteString(mod.Repository.HtmlUrl)
	SendMessage(Message{
		event:   eventName,
		action:  mod.Action,
		repo:    mod.Repository.FullName,
		repoUrl: mod.Repository.HtmlUrl,
		user:    mod.Sender.Login,
		userUrl: mod.Sender.HtmlUrl,
		detail:  detail.String(),
	})
}
