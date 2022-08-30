package config

type OpenAPISpec struct {
	Openapi string `yaml:"openapi" json:"openapi"`
	Info    struct {
		Title       string `yaml:"title" json:"title"`
		Description string `yaml:"description" json:"description"`
		Version     string `yaml:"version" json:"version"`
	} `yaml:"info" json:"info"`
	Paths struct {
		Services struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Responses   struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/services" json:"/services"`
		Timelines struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/timelines" json:"/timelines"`
		Commits struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/commits" json:"/commits"`
		Deploys struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/deploys" json:"/deploys"`
		ServicesService struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/services/{service}" json:"/services/{service}"`
		ServicesServiceTimelines struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/services/{service}/timelines" json:"/services/{service}/timelines"`
		ServicesServiceCommits struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/services/{service}/commits" json:"/services/{service}/commits"`
		ServicesServiceDeploys struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/services/{service}/deploys" json:"/services/{service}/deploys"`
		TimelinesRef struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/timelines/{ref}" json:"/timelines/{ref}"`
		CommitsRef struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/commits/{ref}" json:"/commits/{ref}"`
		DeploysRef struct {
			Get struct {
				Summary     string `yaml:"summary" json:"summary"`
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"parameters" json:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Num400 struct {
						Ref string `yaml:"$ref" json:"$ref"`
					} `yaml:"400" json:"400"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"get" json:"get"`
		} `yaml:"/deploys/{ref}" json:"/deploys/{ref}"`
		GithubWebhook struct {
			Post struct {
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				RequestBody struct {
					Description string `yaml:"description" json:"description"`
					Required    bool   `yaml:"required" json:"required"`
					Content     struct {
						ApplicationJSON struct {
							Schema struct {
								Ref string `yaml:"$ref" json:"$ref"`
							} `yaml:"schema" json:"schema"`
						} `yaml:"application/json" json:"application/json"`
					} `yaml:"content" json:"content"`
				} `yaml:"requestBody" json:"requestBody"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Default struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"default" json:"default"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"post" json:"post"`
		} `yaml:"/github-webhook" json:"/github-webhook"`
		GitlabWebhook struct {
			Post struct {
				Description string `yaml:"description" json:"description"`
				OperationID string `yaml:"operationId" json:"operationId"`
				RequestBody struct {
					Description string `yaml:"description" json:"description"`
					Required    bool   `yaml:"required" json:"required"`
					Content     struct {
						ApplicationJSON struct {
							Schema struct {
								Ref string `yaml:"$ref" json:"$ref"`
							} `yaml:"schema" json:"schema"`
						} `yaml:"application/json" json:"application/json"`
					} `yaml:"content" json:"content"`
				} `yaml:"requestBody" json:"requestBody"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"200" json:"200"`
					Default struct {
						Description string `yaml:"description" json:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref" json:"$ref"`
								} `yaml:"schema" json:"schema"`
							} `yaml:"application/json" json:"application/json"`
						} `yaml:"content" json:"content"`
					} `yaml:"default" json:"default"`
				} `yaml:"responses" json:"responses"`
			} `yaml:"post" json:"post"`
		} `yaml:"/gitlab-webhook" json:"/gitlab-webhook"`
	} `yaml:"paths" json:"paths"`
	Components struct {
		Schemas struct {
			ServicesData struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					Count struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"count" json:"count"`
					Data struct {
						Type  string `yaml:"type" json:"type"`
						Items struct {
							Ref string `yaml:"$ref" json:"$ref"`
						} `yaml:"items" json:"items"`
					} `yaml:"data" json:"data"`
				} `yaml:"properties" json:"properties"`
			} `yaml:"ServicesData" json:"ServicesData"`
			TimelinesData struct {
				Type  string `yaml:"type" json:"type"`
				Items struct {
					Ref string `yaml:"$ref" json:"$ref"`
				} `yaml:"items" json:"items"`
			} `yaml:"TimelinesData" json:"TimelinesData"`
			Data struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					ID struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"id" json:"id"`
					ServiceID struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"service_id" json:"service_id"`
					Timestamp struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"timestamp" json:"timestamp"`
					Type struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"type" json:"type"`
					Repo struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"repo" json:"repo"`
					Ref struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"ref" json:"ref"`
					Author struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"author" json:"author"`
					MergedBy struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"merged_by" json:"merged_by"`
					Message struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"message" json:"message"`
					Namespace struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"namespace" json:"namespace"`
					Cluster struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"cluster" json:"cluster"`
					Image struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"image" json:"image"`
					URL struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"url" json:"url"`
				} `yaml:"properties" json:"properties"`
			} `yaml:"Data" json:"Data"`
			Service struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					ID struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"id" json:"id"`
					Name struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"name" json:"name"`
					DisplayName struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"display_name" json:"display_name"`
					GhRepo struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"gh_repo" json:"gh_repo"`
					GlRepo struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"gl_repo" json:"gl_repo"`
					DeployFile struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"deploy_file" json:"deploy_file"`
					Namespace struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"namespace" json:"namespace"`
					Branch struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"branch" json:"branch"`
					LatestCommit struct {
						Ref      string `yaml:"$ref" json:"$ref"`
						Optional bool   `yaml:"optional" json:"optional"`
					} `yaml:"latest_commit" json:"latest_commit"`
					LatestDeploy struct {
						Ref      string `yaml:"$ref" json:"$ref"`
						Optional bool   `yaml:"optional" json:"optional"`
					} `yaml:"latest_deploy" json:"latest_deploy"`
				} `yaml:"properties" json:"properties"`
			} `yaml:"Service" json:"Service"`
			Timeline struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					Count struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"count" json:"count"`
					Data struct {
						Type  string `yaml:"type" json:"type"`
						Items struct {
							Ref string `yaml:"$ref" json:"$ref"`
						} `yaml:"items" json:"items"`
					} `yaml:"data" json:"data"`
				} `yaml:"properties" json:"properties"`
			} `yaml:"Timeline" json:"Timeline"`
			Error struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					Message struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"message" json:"message"`
				} `yaml:"properties" json:"properties"`
				Required []string `yaml:"required" json:"required"`
			} `yaml:"Error" json:"Error"`
			Message struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					Msg struct {
						Type        string `yaml:"type" json:"type"`
						Description string `yaml:"description" json:"description"`
					} `yaml:"msg" json:"msg"`
				} `yaml:"properties" json:"properties"`
				Required []string `yaml:"required" json:"required"`
			} `yaml:"Message" json:"Message"`
			GithubWebhook struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					Ref struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"ref" json:"ref"`
					Before struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"before" json:"before"`
					After struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"after" json:"after"`
					Repository struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							ID struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"id" json:"id"`
							NodeID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"node_id" json:"node_id"`
							Name struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"name" json:"name"`
							FullName struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"full_name" json:"full_name"`
							Private struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"private" json:"private"`
							Owner struct {
								Type       string `yaml:"type" json:"type"`
								Properties struct {
									Name struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"name" json:"name"`
									Email struct {
										Type   string `yaml:"type" json:"type"`
										Format string `yaml:"format" json:"format"`
									} `yaml:"email" json:"email"`
									Login struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"login" json:"login"`
									ID struct {
										Type   string `yaml:"type" json:"type"`
										Format string `yaml:"format" json:"format"`
									} `yaml:"id" json:"id"`
									NodeID struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"node_id" json:"node_id"`
									AvatarURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"avatar_url" json:"avatar_url"`
									GravatarID struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"gravatar_id" json:"gravatar_id"`
									URL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"url" json:"url"`
									HTMLURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"html_url" json:"html_url"`
									FollowersURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"followers_url" json:"followers_url"`
									FollowingURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"following_url" json:"following_url"`
									GistsURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"gists_url" json:"gists_url"`
									StarredURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"starred_url" json:"starred_url"`
									SubscriptionsURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"subscriptions_url" json:"subscriptions_url"`
									OrganizationsURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"organizations_url" json:"organizations_url"`
									ReposURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"repos_url" json:"repos_url"`
									EventsURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"events_url" json:"events_url"`
									ReceivedEventsURL struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"received_events_url" json:"received_events_url"`
									Type struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"type" json:"type"`
									SiteAdmin struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"site_admin" json:"site_admin"`
								} `yaml:"properties" json:"properties"`
							} `yaml:"owner" json:"owner"`
							HTMLURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"html_url" json:"html_url"`
							Description struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"description" json:"description"`
							Fork struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"fork" json:"fork"`
							URL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"url" json:"url"`
							ForksURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"forks_url" json:"forks_url"`
							KeysURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"keys_url" json:"keys_url"`
							CollaboratorsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"collaborators_url" json:"collaborators_url"`
							TeamsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"teams_url" json:"teams_url"`
							HooksURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"hooks_url" json:"hooks_url"`
							IssueEventsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"issue_events_url" json:"issue_events_url"`
							EventsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"events_url" json:"events_url"`
							AssigneesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"assignees_url" json:"assignees_url"`
							BranchesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"branches_url" json:"branches_url"`
							TagsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"tags_url" json:"tags_url"`
							BlobsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"blobs_url" json:"blobs_url"`
							GitTagsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_tags_url" json:"git_tags_url"`
							GitRefsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_refs_url" json:"git_refs_url"`
							TreesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"trees_url" json:"trees_url"`
							StatusesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"statuses_url" json:"statuses_url"`
							LanguagesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"languages_url" json:"languages_url"`
							StargazersURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"stargazers_url" json:"stargazers_url"`
							ContributorsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"contributors_url" json:"contributors_url"`
							SubscribersURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"subscribers_url" json:"subscribers_url"`
							SubscriptionURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"subscription_url" json:"subscription_url"`
							CommitsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"commits_url" json:"commits_url"`
							GitCommitsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_commits_url" json:"git_commits_url"`
							CommentsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"comments_url" json:"comments_url"`
							IssueCommentURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"issue_comment_url" json:"issue_comment_url"`
							ContentsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"contents_url" json:"contents_url"`
							CompareURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"compare_url" json:"compare_url"`
							MergesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"merges_url" json:"merges_url"`
							ArchiveURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"archive_url" json:"archive_url"`
							DownloadsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"downloads_url" json:"downloads_url"`
							IssuesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"issues_url" json:"issues_url"`
							PullsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"pulls_url" json:"pulls_url"`
							MilestonesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"milestones_url" json:"milestones_url"`
							NotificationsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"notifications_url" json:"notifications_url"`
							LabelsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"labels_url" json:"labels_url"`
							ReleasesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"releases_url" json:"releases_url"`
							DeploymentsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"deployments_url" json:"deployments_url"`
							CreatedAt struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"created_at" json:"created_at"`
							UpdatedAt struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"updated_at" json:"updated_at"`
							PushedAt struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"pushed_at" json:"pushed_at"`
							GitURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_url" json:"git_url"`
							SSHURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"ssh_url" json:"ssh_url"`
							CloneURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"clone_url" json:"clone_url"`
							SvnURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"svn_url" json:"svn_url"`
							Homepage struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"homepage" json:"homepage"`
							Size struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"size" json:"size"`
							StargazersCount struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"stargazers_count" json:"stargazers_count"`
							WatchersCount struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"watchers_count" json:"watchers_count"`
							Language struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"language" json:"language"`
							HasIssues struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"has_issues" json:"has_issues"`
							HasProjects struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"has_projects" json:"has_projects"`
							HasDownloads struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"has_downloads" json:"has_downloads"`
							HasWiki struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"has_wiki" json:"has_wiki"`
							HasPages struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"has_pages" json:"has_pages"`
							ForksCount struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"forks_count" json:"forks_count"`
							MirrorURL struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"mirror_url" json:"mirror_url"`
							Archived struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"archived" json:"archived"`
							Disabled struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"disabled" json:"disabled"`
							OpenIssuesCount struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"open_issues_count" json:"open_issues_count"`
							License struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"license" json:"license"`
							AllowForking struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"allow_forking" json:"allow_forking"`
							IsTemplate struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"is_template" json:"is_template"`
							Topics struct {
								Type  string `yaml:"type" json:"type"`
								Items struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"items" json:"items"`
							} `yaml:"topics" json:"topics"`
							Visibility struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"visibility" json:"visibility"`
							Forks struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"forks" json:"forks"`
							OpenIssues struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"open_issues" json:"open_issues"`
							Watchers struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"watchers" json:"watchers"`
							DefaultBranch struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"default_branch" json:"default_branch"`
							Stargazers struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"stargazers" json:"stargazers"`
							MasterBranch struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"master_branch" json:"master_branch"`
							Organization struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"organization" json:"organization"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"repository" json:"repository"`
					Pusher struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							Name struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"name" json:"name"`
							Email struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"email" json:"email"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"pusher" json:"pusher"`
					Organization struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							Login struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"login" json:"login"`
							ID struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"id" json:"id"`
							NodeID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"node_id" json:"node_id"`
							URL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"url" json:"url"`
							ReposURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"repos_url" json:"repos_url"`
							EventsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"events_url" json:"events_url"`
							HooksURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"hooks_url" json:"hooks_url"`
							IssuesURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"issues_url" json:"issues_url"`
							MembersURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"members_url" json:"members_url"`
							PublicMembersURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"public_members_url" json:"public_members_url"`
							AvatarURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"avatar_url" json:"avatar_url"`
							Description struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"description" json:"description"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"organization" json:"organization"`
					Enterprise struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							ID struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"id" json:"id"`
							Slug struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"slug" json:"slug"`
							Name struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"name" json:"name"`
							NodeID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"node_id" json:"node_id"`
							AvatarURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"avatar_url" json:"avatar_url"`
							Description struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"description" json:"description"`
							WebsiteURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"website_url" json:"website_url"`
							HTMLURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"html_url" json:"html_url"`
							CreatedAt struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"created_at" json:"created_at"`
							UpdatedAt struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"updated_at" json:"updated_at"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"enterprise" json:"enterprise"`
					Sender struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							Login struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"login" json:"login"`
							ID struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"id" json:"id"`
							NodeID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"node_id" json:"node_id"`
							AvatarURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"avatar_url" json:"avatar_url"`
							GravatarID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"gravatar_id" json:"gravatar_id"`
							URL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"url" json:"url"`
							HTMLURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"html_url" json:"html_url"`
							FollowersURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"followers_url" json:"followers_url"`
							FollowingURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"following_url" json:"following_url"`
							GistsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"gists_url" json:"gists_url"`
							StarredURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"starred_url" json:"starred_url"`
							SubscriptionsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"subscriptions_url" json:"subscriptions_url"`
							OrganizationsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"organizations_url" json:"organizations_url"`
							ReposURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"repos_url" json:"repos_url"`
							EventsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"events_url" json:"events_url"`
							ReceivedEventsURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"received_events_url" json:"received_events_url"`
							Type struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"type" json:"type"`
							SiteAdmin struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"site_admin" json:"site_admin"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"sender" json:"sender"`
					Created struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"created" json:"created"`
					Deleted struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"deleted" json:"deleted"`
					Forced struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"forced" json:"forced"`
					BaseRef struct {
						Type   string `yaml:"type" json:"type"`
						Format string `yaml:"format" json:"format"`
					} `yaml:"base_ref" json:"base_ref"`
					Compare struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"compare" json:"compare"`
					Commits struct {
						Type  string `yaml:"type" json:"type"`
						Items struct {
							Type       string `yaml:"type" json:"type"`
							Properties struct {
								ID struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"id" json:"id"`
								TreeID struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"tree_id" json:"tree_id"`
								Distinct struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"distinct" json:"distinct"`
								Message struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"message" json:"message"`
								Timestamp struct {
									Type   string `yaml:"type" json:"type"`
									Format string `yaml:"format" json:"format"`
								} `yaml:"timestamp" json:"timestamp"`
								URL struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"url" json:"url"`
								Author struct {
									Type       string `yaml:"type" json:"type"`
									Properties struct {
										Name struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"name" json:"name"`
										Email struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"email" json:"email"`
										Username struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"username" json:"username"`
									} `yaml:"properties" json:"properties"`
								} `yaml:"author" json:"author"`
								Committer struct {
									Type       string `yaml:"type" json:"type"`
									Properties struct {
										Name struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"name" json:"name"`
										Email struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"email" json:"email"`
										Username struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"username" json:"username"`
									} `yaml:"properties" json:"properties"`
								} `yaml:"committer" json:"committer"`
								Added struct {
									Type  string `yaml:"type" json:"type"`
									Items struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"items" json:"items"`
								} `yaml:"added" json:"added"`
								Removed struct {
									Type  string `yaml:"type" json:"type"`
									Items struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"items" json:"items"`
								} `yaml:"removed" json:"removed"`
								Modified struct {
									Type  string `yaml:"type" json:"type"`
									Items struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"items" json:"items"`
								} `yaml:"modified" json:"modified"`
							} `yaml:"properties" json:"properties"`
						} `yaml:"items" json:"items"`
					} `yaml:"commits" json:"commits"`
					HeadCommit struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							ID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"id" json:"id"`
							TreeID struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"tree_id" json:"tree_id"`
							Distinct struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"distinct" json:"distinct"`
							Message struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"message" json:"message"`
							Timestamp struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"timestamp" json:"timestamp"`
							URL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"url" json:"url"`
							Author struct {
								Type       string `yaml:"type" json:"type"`
								Properties struct {
									Name struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"name" json:"name"`
									Email struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"email" json:"email"`
									Username struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"username" json:"username"`
								} `yaml:"properties" json:"properties"`
							} `yaml:"author" json:"author"`
							Committer struct {
								Type       string `yaml:"type" json:"type"`
								Properties struct {
									Name struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"name" json:"name"`
									Email struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"email" json:"email"`
									Username struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"username" json:"username"`
								} `yaml:"properties" json:"properties"`
							} `yaml:"committer" json:"committer"`
							Added struct {
								Type  string `yaml:"type" json:"type"`
								Items struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"items" json:"items"`
							} `yaml:"added" json:"added"`
							Removed struct {
								Type  string `yaml:"type" json:"type"`
								Items struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"items" json:"items"`
							} `yaml:"removed" json:"removed"`
							Modified struct {
								Type  string `yaml:"type" json:"type"`
								Items struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"items" json:"items"`
							} `yaml:"modified" json:"modified"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"head_commit" json:"head_commit"`
				} `yaml:"properties" json:"properties"`
			} `yaml:"GithubWebhook" json:"GithubWebhook"`
			GitlabWebhook struct {
				Type       string `yaml:"type" json:"type"`
				Properties struct {
					ObjectKind struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"object_kind" json:"object_kind"`
					EventName struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"event_name" json:"event_name"`
					Before struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"before" json:"before"`
					After struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"after" json:"after"`
					Ref struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"ref" json:"ref"`
					CheckoutSha struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"checkout_sha" json:"checkout_sha"`
					Message struct {
						Type   string `yaml:"type" json:"type"`
						Format string `yaml:"format" json:"format"`
					} `yaml:"message" json:"message"`
					UserID struct {
						Type   string `yaml:"type" json:"type"`
						Format string `yaml:"format" json:"format"`
					} `yaml:"user_id" json:"user_id"`
					UserName struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"user_name" json:"user_name"`
					UserUsername struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"user_username" json:"user_username"`
					UserEmail struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"user_email" json:"user_email"`
					UserAvatar struct {
						Type string `yaml:"type" json:"type"`
					} `yaml:"user_avatar" json:"user_avatar"`
					ProjectID struct {
						Type   string `yaml:"type" json:"type"`
						Format string `yaml:"format" json:"format"`
					} `yaml:"project_id" json:"project_id"`
					Project struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							ID struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"id" json:"id"`
							Name struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"name" json:"name"`
							Description struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"description" json:"description"`
							WebURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"web_url" json:"web_url"`
							AvatarURL struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"avatar_url" json:"avatar_url"`
							GitSSHURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_ssh_url" json:"git_ssh_url"`
							GitHTTPURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_http_url" json:"git_http_url"`
							Namespace struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"namespace" json:"namespace"`
							VisibilityLevel struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"visibility_level" json:"visibility_level"`
							PathWithNamespace struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"path_with_namespace" json:"path_with_namespace"`
							DefaultBranch struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"default_branch" json:"default_branch"`
							CiConfigPath struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"ci_config_path" json:"ci_config_path"`
							Homepage struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"homepage" json:"homepage"`
							URL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"url" json:"url"`
							SSHURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"ssh_url" json:"ssh_url"`
							HTTPURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"http_url" json:"http_url"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"project" json:"project"`
					Commits struct {
						Type  string `yaml:"type" json:"type"`
						Items struct {
							Type       string `yaml:"type" json:"type"`
							Properties struct {
								ID struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"id" json:"id"`
								Message struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"message" json:"message"`
								Title struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"title" json:"title"`
								Timestamp struct {
									Type   string `yaml:"type" json:"type"`
									Format string `yaml:"format" json:"format"`
								} `yaml:"timestamp" json:"timestamp"`
								URL struct {
									Type string `yaml:"type" json:"type"`
								} `yaml:"url" json:"url"`
								Author struct {
									Type       string `yaml:"type" json:"type"`
									Properties struct {
										Name struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"name" json:"name"`
										Email struct {
											Type string `yaml:"type" json:"type"`
										} `yaml:"email" json:"email"`
									} `yaml:"properties" json:"properties"`
								} `yaml:"author" json:"author"`
								Added struct {
									Type  string `yaml:"type" json:"type"`
									Items struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"items" json:"items"`
								} `yaml:"added" json:"added"`
								Modified struct {
									Type  string `yaml:"type" json:"type"`
									Items struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"items" json:"items"`
								} `yaml:"modified" json:"modified"`
								Removed struct {
									Type  string `yaml:"type" json:"type"`
									Items struct {
										Type string `yaml:"type" json:"type"`
									} `yaml:"items" json:"items"`
								} `yaml:"removed" json:"removed"`
							} `yaml:"properties" json:"properties"`
						} `yaml:"items" json:"items"`
					} `yaml:"commits" json:"commits"`
					TotalCommitsCount struct {
						Type   string `yaml:"type" json:"type"`
						Format string `yaml:"format" json:"format"`
					} `yaml:"total_commits_count" json:"total_commits_count"`
					Repository struct {
						Type       string `yaml:"type" json:"type"`
						Properties struct {
							Name struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"name" json:"name"`
							URL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"url" json:"url"`
							Description struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"description" json:"description"`
							Homepage struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"homepage" json:"homepage"`
							GitHTTPURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_http_url" json:"git_http_url"`
							GitSSHURL struct {
								Type string `yaml:"type" json:"type"`
							} `yaml:"git_ssh_url" json:"git_ssh_url"`
							VisibilityLevel struct {
								Type   string `yaml:"type" json:"type"`
								Format string `yaml:"format" json:"format"`
							} `yaml:"visibility_level" json:"visibility_level"`
						} `yaml:"properties" json:"properties"`
					} `yaml:"repository" json:"repository"`
				} `yaml:"properties" json:"properties"`
			} `yaml:"GitlabWebhook" json:"GitlabWebhook"`
		} `yaml:"schemas" json:"schemas"`
		Parameters struct {
			Service struct {
				In       string `yaml:"in" json:"in"`
				Name     string `yaml:"name" json:"name"`
				Required bool   `yaml:"required" json:"required"`
				Schema   struct {
					Type string `yaml:"type" json:"type"`
				} `yaml:"schema" json:"schema"`
				Description string `yaml:"description" json:"description"`
			} `yaml:"service" json:"service"`
			Ref struct {
				In       string `yaml:"in" json:"in"`
				Name     string `yaml:"name" json:"name"`
				Required bool   `yaml:"required" json:"required"`
				Schema   struct {
					Type string `yaml:"type" json:"type"`
				} `yaml:"schema" json:"schema"`
				Description string `yaml:"description" json:"description"`
			} `yaml:"ref" json:"ref"`
			OffsetParam struct {
				In       string `yaml:"in" json:"in"`
				Name     string `yaml:"name" json:"name"`
				Required bool   `yaml:"required" json:"required"`
				Schema   struct {
					Type    string `yaml:"type" json:"type"`
					Minimum int    `yaml:"minimum" json:"minimum"`
					Default int    `yaml:"default" json:"default"`
				} `yaml:"schema" json:"schema"`
				Description string `yaml:"description" json:"description"`
			} `yaml:"offsetParam" json:"offsetParam"`
			LimitParam struct {
				In       string `yaml:"in" json:"in"`
				Name     string `yaml:"name" json:"name"`
				Required bool   `yaml:"required" json:"required"`
				Schema   struct {
					Type    string `yaml:"type" json:"type"`
					Minimum int    `yaml:"minimum" json:"minimum"`
					Maximum int    `yaml:"maximum" json:"maximum"`
					Default int    `yaml:"default" json:"default"`
				} `yaml:"schema" json:"schema"`
				Description string `yaml:"description" json:"description"`
			} `yaml:"limitParam" json:"limitParam"`
		} `yaml:"parameters" json:"parameters"`
		Responses struct {
			BadRequest struct {
				Description string `yaml:"description" json:"description"`
				Content     struct {
					ApplicationJSON struct {
						Schema struct {
							Ref string `yaml:"$ref" json:"$ref"`
						} `yaml:"schema" json:"schema"`
					} `yaml:"application/json" json:"application/json"`
				} `yaml:"content" json:"content"`
			} `yaml:"BadRequest" json:"BadRequest"`
		} `yaml:"responses" json:"responses"`
	} `yaml:"components" json:"components"`
}
