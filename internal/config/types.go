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
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/services"`
		Timelines struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/timelines"`
		Commits struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/commits"`
		Deploys struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/deploys"`
		ServicesService struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/services/{service}"`
		ServicesServiceTimelines struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/services/{service}/timelines"`
		ServicesServiceCommits struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/services/{service}/commits"`
		ServicesServiceDeploys struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/services/{service}/deploys"`
		TimelinesRef struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/timelines/{ref}"`
		CommitsRef struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/commits/{ref}"`
		DeploysRef struct {
			Get struct {
				Summary     string `yaml:"summary"`
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				Parameters  []struct {
					Ref string `yaml:"$ref"`
				} `yaml:"parameters"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Num400 struct {
						Ref string `yaml:"$ref"`
					} `yaml:"400"`
				} `yaml:"responses"`
			} `yaml:"get"`
		} `yaml:"/deploys/{ref}"`
		GithubWebhook struct {
			Post struct {
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				RequestBody struct {
					Description string `yaml:"description"`
					Required    bool   `yaml:"required"`
					Content     struct {
						ApplicationJSON struct {
							Schema struct {
								Ref string `yaml:"$ref"`
							} `yaml:"schema"`
						} `yaml:"application/json"`
					} `yaml:"content"`
				} `yaml:"requestBody"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Default struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"default"`
				} `yaml:"responses"`
			} `yaml:"post"`
		} `yaml:"/github-webhook"`
		GitlabWebhook struct {
			Post struct {
				Description string `yaml:"description"`
				OperationID string `yaml:"operationId"`
				RequestBody struct {
					Description string `yaml:"description"`
					Required    bool   `yaml:"required"`
					Content     struct {
						ApplicationJSON struct {
							Schema struct {
								Ref string `yaml:"$ref"`
							} `yaml:"schema"`
						} `yaml:"application/json"`
					} `yaml:"content"`
				} `yaml:"requestBody"`
				Responses struct {
					Num200 struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"200"`
					Default struct {
						Description string `yaml:"description"`
						Content     struct {
							ApplicationJSON struct {
								Schema struct {
									Ref string `yaml:"$ref"`
								} `yaml:"schema"`
							} `yaml:"application/json"`
						} `yaml:"content"`
					} `yaml:"default"`
				} `yaml:"responses"`
			} `yaml:"post"`
		} `yaml:"/gitlab-webhook"`
	} `yaml:"paths"`
	Components struct {
		Schemas struct {
			ServicesData struct {
				Type       string `yaml:"type"`
				Properties struct {
					Count struct {
						Type string `yaml:"type"`
					} `yaml:"count"`
					Data struct {
						Type  string `yaml:"type"`
						Items struct {
							Ref string `yaml:"$ref"`
						} `yaml:"items"`
					} `yaml:"data"`
				} `yaml:"properties"`
			} `yaml:"ServicesData"`
			TimelinesData struct {
				Type       string `yaml:"type"`
				Properties struct {
					Count struct {
						Type string `yaml:"type"`
					} `yaml:"count"`
					Data struct {
						Type  string `yaml:"type"`
						Items struct {
							Ref string `yaml:"$ref"`
						} `yaml:"items"`
					} `yaml:"data"`
				} `yaml:"properties"`
			} `yaml:"TimelinesData"`
			Data struct {
				Type       string `yaml:"type"`
				Properties struct {
					ID struct {
						Type string `yaml:"type"`
					} `yaml:"id"`
					ServiceID struct {
						Type string `yaml:"type"`
					} `yaml:"service_id"`
					Timestamp struct {
						Type string `yaml:"type"`
					} `yaml:"timestamp"`
					Type struct {
						Type string `yaml:"type"`
					} `yaml:"type"`
					Repo struct {
						Type string `yaml:"type"`
					} `yaml:"repo"`
					Ref struct {
						Type string `yaml:"type"`
					} `yaml:"ref"`
					Author struct {
						Type string `yaml:"type"`
					} `yaml:"author"`
					MergedBy struct {
						Type string `yaml:"type"`
					} `yaml:"merged_by"`
					Message struct {
						Type string `yaml:"type"`
					} `yaml:"message"`
					Namespace struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"namespace"`
					Cluster struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"cluster"`
					Image struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"image"`
					URL struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"url"`
				} `yaml:"properties"`
			} `yaml:"Data"`
			Service struct {
				Type       string `yaml:"type"`
				Properties struct {
					ID struct {
						Type string `yaml:"type"`
					} `yaml:"id"`
					Name struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"name"`
					DisplayName struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"display_name"`
					Tenant struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"tenant"`
					GhRepo struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"gh_repo"`
					GlRepo struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"gl_repo"`
					DeployFile struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"deploy_file"`
					Namespace struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"namespace"`
					Branch struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"branch"`
					LatestCommit struct {
						Ref      string `yaml:"$ref"`
						Optional bool   `yaml:"optional"`
					} `yaml:"latest_commit"`
					LatestDeploy struct {
						Ref      string `yaml:"$ref"`
						Optional bool   `yaml:"optional"`
					} `yaml:"latest_deploy"`
				} `yaml:"properties"`
			} `yaml:"Service"`
			Timeline struct {
				Type       string `yaml:"type"`
				Properties struct {
					Count struct {
						Type string `yaml:"type"`
					} `yaml:"count"`
					Data struct {
						Type  string `yaml:"type"`
						Items struct {
							Ref string `yaml:"$ref"`
						} `yaml:"items"`
					} `yaml:"data"`
				} `yaml:"properties"`
			} `yaml:"Timeline"`
			Error struct {
				Type       string `yaml:"type"`
				Properties struct {
					Message struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"message"`
				} `yaml:"properties"`
				Required []string `yaml:"required"`
			} `yaml:"Error"`
			Message struct {
				Type       string `yaml:"type"`
				Properties struct {
					Msg struct {
						Type        string `yaml:"type"`
						Description string `yaml:"description"`
					} `yaml:"msg"`
				} `yaml:"properties"`
				Required []string `yaml:"required"`
			} `yaml:"Message"`
			GithubWebhook struct {
				Type       string `yaml:"type"`
				Properties struct {
					Ref struct {
						Type string `yaml:"type"`
					} `yaml:"ref"`
					Before struct {
						Type string `yaml:"type"`
					} `yaml:"before"`
					After struct {
						Type string `yaml:"type"`
					} `yaml:"after"`
					Repository struct {
						Type       string `yaml:"type"`
						Properties struct {
							ID struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"id"`
							NodeID struct {
								Type string `yaml:"type"`
							} `yaml:"node_id"`
							Name struct {
								Type string `yaml:"type"`
							} `yaml:"name"`
							FullName struct {
								Type string `yaml:"type"`
							} `yaml:"full_name"`
							Private struct {
								Type string `yaml:"type"`
							} `yaml:"private"`
							Owner struct {
								Type       string `yaml:"type"`
								Properties struct {
									Name struct {
										Type string `yaml:"type"`
									} `yaml:"name"`
									Email struct {
										Type   string `yaml:"type"`
										Format string `yaml:"format"`
									} `yaml:"email"`
									Login struct {
										Type string `yaml:"type"`
									} `yaml:"login"`
									ID struct {
										Type   string `yaml:"type"`
										Format string `yaml:"format"`
									} `yaml:"id"`
									NodeID struct {
										Type string `yaml:"type"`
									} `yaml:"node_id"`
									AvatarURL struct {
										Type string `yaml:"type"`
									} `yaml:"avatar_url"`
									GravatarID struct {
										Type string `yaml:"type"`
									} `yaml:"gravatar_id"`
									URL struct {
										Type string `yaml:"type"`
									} `yaml:"url"`
									HTMLURL struct {
										Type string `yaml:"type"`
									} `yaml:"html_url"`
									FollowersURL struct {
										Type string `yaml:"type"`
									} `yaml:"followers_url"`
									FollowingURL struct {
										Type string `yaml:"type"`
									} `yaml:"following_url"`
									GistsURL struct {
										Type string `yaml:"type"`
									} `yaml:"gists_url"`
									StarredURL struct {
										Type string `yaml:"type"`
									} `yaml:"starred_url"`
									SubscriptionsURL struct {
										Type string `yaml:"type"`
									} `yaml:"subscriptions_url"`
									OrganizationsURL struct {
										Type string `yaml:"type"`
									} `yaml:"organizations_url"`
									ReposURL struct {
										Type string `yaml:"type"`
									} `yaml:"repos_url"`
									EventsURL struct {
										Type string `yaml:"type"`
									} `yaml:"events_url"`
									ReceivedEventsURL struct {
										Type string `yaml:"type"`
									} `yaml:"received_events_url"`
									Type struct {
										Type string `yaml:"type"`
									} `yaml:"type"`
									SiteAdmin struct {
										Type string `yaml:"type"`
									} `yaml:"site_admin"`
								} `yaml:"properties"`
							} `yaml:"owner"`
							HTMLURL struct {
								Type string `yaml:"type"`
							} `yaml:"html_url"`
							Description struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"description"`
							Fork struct {
								Type string `yaml:"type"`
							} `yaml:"fork"`
							URL struct {
								Type string `yaml:"type"`
							} `yaml:"url"`
							ForksURL struct {
								Type string `yaml:"type"`
							} `yaml:"forks_url"`
							KeysURL struct {
								Type string `yaml:"type"`
							} `yaml:"keys_url"`
							CollaboratorsURL struct {
								Type string `yaml:"type"`
							} `yaml:"collaborators_url"`
							TeamsURL struct {
								Type string `yaml:"type"`
							} `yaml:"teams_url"`
							HooksURL struct {
								Type string `yaml:"type"`
							} `yaml:"hooks_url"`
							IssueEventsURL struct {
								Type string `yaml:"type"`
							} `yaml:"issue_events_url"`
							EventsURL struct {
								Type string `yaml:"type"`
							} `yaml:"events_url"`
							AssigneesURL struct {
								Type string `yaml:"type"`
							} `yaml:"assignees_url"`
							BranchesURL struct {
								Type string `yaml:"type"`
							} `yaml:"branches_url"`
							TagsURL struct {
								Type string `yaml:"type"`
							} `yaml:"tags_url"`
							BlobsURL struct {
								Type string `yaml:"type"`
							} `yaml:"blobs_url"`
							GitTagsURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_tags_url"`
							GitRefsURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_refs_url"`
							TreesURL struct {
								Type string `yaml:"type"`
							} `yaml:"trees_url"`
							StatusesURL struct {
								Type string `yaml:"type"`
							} `yaml:"statuses_url"`
							LanguagesURL struct {
								Type string `yaml:"type"`
							} `yaml:"languages_url"`
							StargazersURL struct {
								Type string `yaml:"type"`
							} `yaml:"stargazers_url"`
							ContributorsURL struct {
								Type string `yaml:"type"`
							} `yaml:"contributors_url"`
							SubscribersURL struct {
								Type string `yaml:"type"`
							} `yaml:"subscribers_url"`
							SubscriptionURL struct {
								Type string `yaml:"type"`
							} `yaml:"subscription_url"`
							CommitsURL struct {
								Type string `yaml:"type"`
							} `yaml:"commits_url"`
							GitCommitsURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_commits_url"`
							CommentsURL struct {
								Type string `yaml:"type"`
							} `yaml:"comments_url"`
							IssueCommentURL struct {
								Type string `yaml:"type"`
							} `yaml:"issue_comment_url"`
							ContentsURL struct {
								Type string `yaml:"type"`
							} `yaml:"contents_url"`
							CompareURL struct {
								Type string `yaml:"type"`
							} `yaml:"compare_url"`
							MergesURL struct {
								Type string `yaml:"type"`
							} `yaml:"merges_url"`
							ArchiveURL struct {
								Type string `yaml:"type"`
							} `yaml:"archive_url"`
							DownloadsURL struct {
								Type string `yaml:"type"`
							} `yaml:"downloads_url"`
							IssuesURL struct {
								Type string `yaml:"type"`
							} `yaml:"issues_url"`
							PullsURL struct {
								Type string `yaml:"type"`
							} `yaml:"pulls_url"`
							MilestonesURL struct {
								Type string `yaml:"type"`
							} `yaml:"milestones_url"`
							NotificationsURL struct {
								Type string `yaml:"type"`
							} `yaml:"notifications_url"`
							LabelsURL struct {
								Type string `yaml:"type"`
							} `yaml:"labels_url"`
							ReleasesURL struct {
								Type string `yaml:"type"`
							} `yaml:"releases_url"`
							DeploymentsURL struct {
								Type string `yaml:"type"`
							} `yaml:"deployments_url"`
							CreatedAt struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"created_at"`
							UpdatedAt struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"updated_at"`
							PushedAt struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"pushed_at"`
							GitURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_url"`
							SSHURL struct {
								Type string `yaml:"type"`
							} `yaml:"ssh_url"`
							CloneURL struct {
								Type string `yaml:"type"`
							} `yaml:"clone_url"`
							SvnURL struct {
								Type string `yaml:"type"`
							} `yaml:"svn_url"`
							Homepage struct {
								Type string `yaml:"type"`
							} `yaml:"homepage"`
							Size struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"size"`
							StargazersCount struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"stargazers_count"`
							WatchersCount struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"watchers_count"`
							Language struct {
								Type string `yaml:"type"`
							} `yaml:"language"`
							HasIssues struct {
								Type string `yaml:"type"`
							} `yaml:"has_issues"`
							HasProjects struct {
								Type string `yaml:"type"`
							} `yaml:"has_projects"`
							HasDownloads struct {
								Type string `yaml:"type"`
							} `yaml:"has_downloads"`
							HasWiki struct {
								Type string `yaml:"type"`
							} `yaml:"has_wiki"`
							HasPages struct {
								Type string `yaml:"type"`
							} `yaml:"has_pages"`
							ForksCount struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"forks_count"`
							MirrorURL struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"mirror_url"`
							Archived struct {
								Type string `yaml:"type"`
							} `yaml:"archived"`
							Disabled struct {
								Type string `yaml:"type"`
							} `yaml:"disabled"`
							OpenIssuesCount struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"open_issues_count"`
							License struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"license"`
							AllowForking struct {
								Type string `yaml:"type"`
							} `yaml:"allow_forking"`
							IsTemplate struct {
								Type string `yaml:"type"`
							} `yaml:"is_template"`
							Topics struct {
								Type  string `yaml:"type"`
								Items struct {
									Type string `yaml:"type"`
								} `yaml:"items"`
							} `yaml:"topics"`
							Visibility struct {
								Type string `yaml:"type"`
							} `yaml:"visibility"`
							Forks struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"forks"`
							OpenIssues struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"open_issues"`
							Watchers struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"watchers"`
							DefaultBranch struct {
								Type string `yaml:"type"`
							} `yaml:"default_branch"`
							Stargazers struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"stargazers"`
							MasterBranch struct {
								Type string `yaml:"type"`
							} `yaml:"master_branch"`
							Organization struct {
								Type string `yaml:"type"`
							} `yaml:"organization"`
						} `yaml:"properties"`
					} `yaml:"repository"`
					Pusher struct {
						Type       string `yaml:"type"`
						Properties struct {
							Name struct {
								Type string `yaml:"type"`
							} `yaml:"name"`
							Email struct {
								Type string `yaml:"type"`
							} `yaml:"email"`
						} `yaml:"properties"`
					} `yaml:"pusher"`
					Organization struct {
						Type       string `yaml:"type"`
						Properties struct {
							Login struct {
								Type string `yaml:"type"`
							} `yaml:"login"`
							ID struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"id"`
							NodeID struct {
								Type string `yaml:"type"`
							} `yaml:"node_id"`
							URL struct {
								Type string `yaml:"type"`
							} `yaml:"url"`
							ReposURL struct {
								Type string `yaml:"type"`
							} `yaml:"repos_url"`
							EventsURL struct {
								Type string `yaml:"type"`
							} `yaml:"events_url"`
							HooksURL struct {
								Type string `yaml:"type"`
							} `yaml:"hooks_url"`
							IssuesURL struct {
								Type string `yaml:"type"`
							} `yaml:"issues_url"`
							MembersURL struct {
								Type string `yaml:"type"`
							} `yaml:"members_url"`
							PublicMembersURL struct {
								Type string `yaml:"type"`
							} `yaml:"public_members_url"`
							AvatarURL struct {
								Type string `yaml:"type"`
							} `yaml:"avatar_url"`
							Description struct {
								Type string `yaml:"type"`
							} `yaml:"description"`
						} `yaml:"properties"`
					} `yaml:"organization"`
					Enterprise struct {
						Type       string `yaml:"type"`
						Properties struct {
							ID struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"id"`
							Slug struct {
								Type string `yaml:"type"`
							} `yaml:"slug"`
							Name struct {
								Type string `yaml:"type"`
							} `yaml:"name"`
							NodeID struct {
								Type string `yaml:"type"`
							} `yaml:"node_id"`
							AvatarURL struct {
								Type string `yaml:"type"`
							} `yaml:"avatar_url"`
							Description struct {
								Type string `yaml:"type"`
							} `yaml:"description"`
							WebsiteURL struct {
								Type string `yaml:"type"`
							} `yaml:"website_url"`
							HTMLURL struct {
								Type string `yaml:"type"`
							} `yaml:"html_url"`
							CreatedAt struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"created_at"`
							UpdatedAt struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"updated_at"`
						} `yaml:"properties"`
					} `yaml:"enterprise"`
					Sender struct {
						Type       string `yaml:"type"`
						Properties struct {
							Login struct {
								Type string `yaml:"type"`
							} `yaml:"login"`
							ID struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"id"`
							NodeID struct {
								Type string `yaml:"type"`
							} `yaml:"node_id"`
							AvatarURL struct {
								Type string `yaml:"type"`
							} `yaml:"avatar_url"`
							GravatarID struct {
								Type string `yaml:"type"`
							} `yaml:"gravatar_id"`
							URL struct {
								Type string `yaml:"type"`
							} `yaml:"url"`
							HTMLURL struct {
								Type string `yaml:"type"`
							} `yaml:"html_url"`
							FollowersURL struct {
								Type string `yaml:"type"`
							} `yaml:"followers_url"`
							FollowingURL struct {
								Type string `yaml:"type"`
							} `yaml:"following_url"`
							GistsURL struct {
								Type string `yaml:"type"`
							} `yaml:"gists_url"`
							StarredURL struct {
								Type string `yaml:"type"`
							} `yaml:"starred_url"`
							SubscriptionsURL struct {
								Type string `yaml:"type"`
							} `yaml:"subscriptions_url"`
							OrganizationsURL struct {
								Type string `yaml:"type"`
							} `yaml:"organizations_url"`
							ReposURL struct {
								Type string `yaml:"type"`
							} `yaml:"repos_url"`
							EventsURL struct {
								Type string `yaml:"type"`
							} `yaml:"events_url"`
							ReceivedEventsURL struct {
								Type string `yaml:"type"`
							} `yaml:"received_events_url"`
							Type struct {
								Type string `yaml:"type"`
							} `yaml:"type"`
							SiteAdmin struct {
								Type string `yaml:"type"`
							} `yaml:"site_admin"`
						} `yaml:"properties"`
					} `yaml:"sender"`
					Created struct {
						Type string `yaml:"type"`
					} `yaml:"created"`
					Deleted struct {
						Type string `yaml:"type"`
					} `yaml:"deleted"`
					Forced struct {
						Type string `yaml:"type"`
					} `yaml:"forced"`
					BaseRef struct {
						Type   string `yaml:"type"`
						Format string `yaml:"format"`
					} `yaml:"base_ref"`
					Compare struct {
						Type string `yaml:"type"`
					} `yaml:"compare"`
					Commits struct {
						Type  string `yaml:"type"`
						Items struct {
							Type       string `yaml:"type"`
							Properties struct {
								ID struct {
									Type string `yaml:"type"`
								} `yaml:"id"`
								TreeID struct {
									Type string `yaml:"type"`
								} `yaml:"tree_id"`
								Distinct struct {
									Type string `yaml:"type"`
								} `yaml:"distinct"`
								Message struct {
									Type string `yaml:"type"`
								} `yaml:"message"`
								Timestamp struct {
									Type   string `yaml:"type"`
									Format string `yaml:"format"`
								} `yaml:"timestamp"`
								URL struct {
									Type string `yaml:"type"`
								} `yaml:"url"`
								Author struct {
									Type       string `yaml:"type"`
									Properties struct {
										Name struct {
											Type string `yaml:"type"`
										} `yaml:"name"`
										Email struct {
											Type string `yaml:"type"`
										} `yaml:"email"`
										Username struct {
											Type string `yaml:"type"`
										} `yaml:"username"`
									} `yaml:"properties"`
								} `yaml:"author"`
								Committer struct {
									Type       string `yaml:"type"`
									Properties struct {
										Name struct {
											Type string `yaml:"type"`
										} `yaml:"name"`
										Email struct {
											Type string `yaml:"type"`
										} `yaml:"email"`
										Username struct {
											Type string `yaml:"type"`
										} `yaml:"username"`
									} `yaml:"properties"`
								} `yaml:"committer"`
								Added struct {
									Type  string `yaml:"type"`
									Items struct {
										Type string `yaml:"type"`
									} `yaml:"items"`
								} `yaml:"added"`
								Removed struct {
									Type  string `yaml:"type"`
									Items struct {
										Type string `yaml:"type"`
									} `yaml:"items"`
								} `yaml:"removed"`
								Modified struct {
									Type  string `yaml:"type"`
									Items struct {
										Type string `yaml:"type"`
									} `yaml:"items"`
								} `yaml:"modified"`
							} `yaml:"properties"`
						} `yaml:"items"`
					} `yaml:"commits"`
					HeadCommit struct {
						Type       string `yaml:"type"`
						Properties struct {
							ID struct {
								Type string `yaml:"type"`
							} `yaml:"id"`
							TreeID struct {
								Type string `yaml:"type"`
							} `yaml:"tree_id"`
							Distinct struct {
								Type string `yaml:"type"`
							} `yaml:"distinct"`
							Message struct {
								Type string `yaml:"type"`
							} `yaml:"message"`
							Timestamp struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"timestamp"`
							URL struct {
								Type string `yaml:"type"`
							} `yaml:"url"`
							Author struct {
								Type       string `yaml:"type"`
								Properties struct {
									Name struct {
										Type string `yaml:"type"`
									} `yaml:"name"`
									Email struct {
										Type string `yaml:"type"`
									} `yaml:"email"`
									Username struct {
										Type string `yaml:"type"`
									} `yaml:"username"`
								} `yaml:"properties"`
							} `yaml:"author"`
							Committer struct {
								Type       string `yaml:"type"`
								Properties struct {
									Name struct {
										Type string `yaml:"type"`
									} `yaml:"name"`
									Email struct {
										Type string `yaml:"type"`
									} `yaml:"email"`
									Username struct {
										Type string `yaml:"type"`
									} `yaml:"username"`
								} `yaml:"properties"`
							} `yaml:"committer"`
							Added struct {
								Type  string `yaml:"type"`
								Items struct {
									Type string `yaml:"type"`
								} `yaml:"items"`
							} `yaml:"added"`
							Removed struct {
								Type  string `yaml:"type"`
								Items struct {
									Type string `yaml:"type"`
								} `yaml:"items"`
							} `yaml:"removed"`
							Modified struct {
								Type  string `yaml:"type"`
								Items struct {
									Type string `yaml:"type"`
								} `yaml:"items"`
							} `yaml:"modified"`
						} `yaml:"properties"`
					} `yaml:"head_commit"`
				} `yaml:"properties"`
			} `yaml:"GithubWebhook"`
			GitlabWebhook struct {
				Type       string `yaml:"type"`
				Properties struct {
					ObjectKind struct {
						Type string `yaml:"type"`
					} `yaml:"object_kind"`
					EventName struct {
						Type string `yaml:"type"`
					} `yaml:"event_name"`
					Before struct {
						Type string `yaml:"type"`
					} `yaml:"before"`
					After struct {
						Type string `yaml:"type"`
					} `yaml:"after"`
					Ref struct {
						Type string `yaml:"type"`
					} `yaml:"ref"`
					CheckoutSha struct {
						Type string `yaml:"type"`
					} `yaml:"checkout_sha"`
					Message struct {
						Type   string `yaml:"type"`
						Format string `yaml:"format"`
					} `yaml:"message"`
					UserID struct {
						Type   string `yaml:"type"`
						Format string `yaml:"format"`
					} `yaml:"user_id"`
					UserName struct {
						Type string `yaml:"type"`
					} `yaml:"user_name"`
					UserUsername struct {
						Type string `yaml:"type"`
					} `yaml:"user_username"`
					UserEmail struct {
						Type string `yaml:"type"`
					} `yaml:"user_email"`
					UserAvatar struct {
						Type string `yaml:"type"`
					} `yaml:"user_avatar"`
					ProjectID struct {
						Type   string `yaml:"type"`
						Format string `yaml:"format"`
					} `yaml:"project_id"`
					Project struct {
						Type       string `yaml:"type"`
						Properties struct {
							ID struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"id"`
							Name struct {
								Type string `yaml:"type"`
							} `yaml:"name"`
							Description struct {
								Type string `yaml:"type"`
							} `yaml:"description"`
							WebURL struct {
								Type string `yaml:"type"`
							} `yaml:"web_url"`
							AvatarURL struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"avatar_url"`
							GitSSHURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_ssh_url"`
							GitHTTPURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_http_url"`
							Namespace struct {
								Type string `yaml:"type"`
							} `yaml:"namespace"`
							VisibilityLevel struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"visibility_level"`
							PathWithNamespace struct {
								Type string `yaml:"type"`
							} `yaml:"path_with_namespace"`
							DefaultBranch struct {
								Type string `yaml:"type"`
							} `yaml:"default_branch"`
							CiConfigPath struct {
								Type string `yaml:"type"`
							} `yaml:"ci_config_path"`
							Homepage struct {
								Type string `yaml:"type"`
							} `yaml:"homepage"`
							URL struct {
								Type string `yaml:"type"`
							} `yaml:"url"`
							SSHURL struct {
								Type string `yaml:"type"`
							} `yaml:"ssh_url"`
							HTTPURL struct {
								Type string `yaml:"type"`
							} `yaml:"http_url"`
						} `yaml:"properties"`
					} `yaml:"project"`
					Commits struct {
						Type  string `yaml:"type"`
						Items struct {
							Type       string `yaml:"type"`
							Properties struct {
								ID struct {
									Type string `yaml:"type"`
								} `yaml:"id"`
								Message struct {
									Type string `yaml:"type"`
								} `yaml:"message"`
								Title struct {
									Type string `yaml:"type"`
								} `yaml:"title"`
								Timestamp struct {
									Type   string `yaml:"type"`
									Format string `yaml:"format"`
								} `yaml:"timestamp"`
								URL struct {
									Type string `yaml:"type"`
								} `yaml:"url"`
								Author struct {
									Type       string `yaml:"type"`
									Properties struct {
										Name struct {
											Type string `yaml:"type"`
										} `yaml:"name"`
										Email struct {
											Type string `yaml:"type"`
										} `yaml:"email"`
									} `yaml:"properties"`
								} `yaml:"author"`
								Added struct {
									Type  string `yaml:"type"`
									Items struct {
										Type string `yaml:"type"`
									} `yaml:"items"`
								} `yaml:"added"`
								Modified struct {
									Type  string `yaml:"type"`
									Items struct {
										Type string `yaml:"type"`
									} `yaml:"items"`
								} `yaml:"modified"`
								Removed struct {
									Type  string `yaml:"type"`
									Items struct {
										Type string `yaml:"type"`
									} `yaml:"items"`
								} `yaml:"removed"`
							} `yaml:"properties"`
						} `yaml:"items"`
					} `yaml:"commits"`
					TotalCommitsCount struct {
						Type   string `yaml:"type"`
						Format string `yaml:"format"`
					} `yaml:"total_commits_count"`
					Repository struct {
						Type       string `yaml:"type"`
						Properties struct {
							Name struct {
								Type string `yaml:"type"`
							} `yaml:"name"`
							URL struct {
								Type string `yaml:"type"`
							} `yaml:"url"`
							Description struct {
								Type string `yaml:"type"`
							} `yaml:"description"`
							Homepage struct {
								Type string `yaml:"type"`
							} `yaml:"homepage"`
							GitHTTPURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_http_url"`
							GitSSHURL struct {
								Type string `yaml:"type"`
							} `yaml:"git_ssh_url"`
							VisibilityLevel struct {
								Type   string `yaml:"type"`
								Format string `yaml:"format"`
							} `yaml:"visibility_level"`
						} `yaml:"properties"`
					} `yaml:"repository"`
				} `yaml:"properties"`
			} `yaml:"GitlabWebhook"`
		} `yaml:"schemas"`
		Parameters struct {
			Service struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"service"`
			Ref struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"ref"`
			OffsetParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type    string `yaml:"type"`
					Minimum int    `yaml:"minimum"`
					Default int    `yaml:"default"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"offsetParam"`
			LimitParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type    string `yaml:"type"`
					Minimum int    `yaml:"minimum"`
					Maximum int    `yaml:"maximum"`
					Default int    `yaml:"default"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"limitParam"`
			RefFilter struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"refFilter"`
			RepoFilter struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"repoFilter"`
			AuthorFilter struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"authorFilter"`
			MergedByFilter struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"mergedByFilter"`
			ClusterFilter struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"clusterFilter"`
			ImageFilter struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"imageFilter"`
			StartDate struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type   string `yaml:"type"`
					Format string `yaml:"format"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"startDate"`
			EndDate struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type   string `yaml:"type"`
					Format string `yaml:"format"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"endDate"`
			NameParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"nameParam"`
			DisplayNameParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"displayNameParam"`
			TenantParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"tenantParam"`
			NamespaceParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"namespaceParam"`
			BranchParam struct {
				In       string `yaml:"in"`
				Name     string `yaml:"name"`
				Required bool   `yaml:"required"`
				Schema   struct {
					Type string `yaml:"type"`
				} `yaml:"schema"`
				Description string `yaml:"description"`
			} `yaml:"branchParam"`
		} `yaml:"parameters"`
		Responses struct {
			BadRequest struct {
				Description string `yaml:"description"`
				Content     struct {
					ApplicationJSON struct {
						Schema struct {
							Ref string `yaml:"$ref"`
						} `yaml:"schema"`
					} `yaml:"application/json"`
				} `yaml:"content"`
			} `yaml:"BadRequest"`
		} `yaml:"responses"`
	} `yaml:"components"`
}
