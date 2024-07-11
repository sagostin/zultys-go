package lib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type SystemSettings struct {
	Command        string `json:"command"`
	Index          int    `json:"index"`
	Success        bool   `json:"success"`
	SystemSettings struct {
		CloudMediaServer struct {
			Description string `json:"description"`
			Options     struct {
				Address struct {
					Value string `json:"value"`
				} `json:"address"`
			} `json:"options"`
		} `json:"cloudMediaServer"`
		CloudServicesSettings struct {
			Description string `json:"description"`
			Options     struct {
				AsrServiceList struct {
					Value []struct {
						ApiKey struct {
							Value   string `json:"value"`
							Enabled bool   `json:"enabled,omitempty"`
						} `json:"apiKey"`
						EndPointUrl struct {
							Value string `json:"value"`
						} `json:"endPointUrl"`
						Model struct {
							Enabled bool   `json:"enabled,omitempty"`
							Value   string `json:"value"`
						} `json:"model"`
						Password struct {
							Enabled bool   `json:"enabled,omitempty"`
							Value   string `json:"value"`
						} `json:"password"`
						ServiceName struct {
							Value string `json:"value"`
						} `json:"serviceName"`
						User struct {
							Enabled bool   `json:"enabled,omitempty"`
							Value   string `json:"value"`
						} `json:"user"`
					} `json:"value"`
				} `json:"asrServiceList"`
				Enabled struct {
					Value bool `json:"value"`
				} `json:"enabled"`
				ServiceName struct {
					Value string `json:"value"`
				} `json:"serviceName"`
			} `json:"options"`
		} `json:"cloudServicesSettings"`
		CompanySettings struct {
			Description string `json:"description"`
			Options     struct {
				CallProgressTones struct {
					Value string `json:"value"`
				} `json:"callProgressTones"`
				CityOrTown struct {
					Value string `json:"value"`
				} `json:"cityOrTown"`
				CompanyName struct {
					Value string `json:"value"`
				} `json:"companyName"`
				CountryCode struct {
					Value int `json:"value"`
				} `json:"countryCode"`
				CountryOrRegion struct {
					Value string `json:"value"`
				} `json:"countryOrRegion"`
				DefaultCodec struct {
					Enabled bool `json:"enabled"`
					Value   int  `json:"value"`
					Visible bool `json:"visible"`
				} `json:"defaultCodec"`
				DefaultDomain struct {
					Value string `json:"value"`
				} `json:"defaultDomain"`
				Language struct {
					Value int `json:"value"`
				} `json:"language"`
				MainPhoneNumber struct {
					Value string `json:"value"`
				} `json:"mainPhoneNumber"`
				NameForThisLocation struct {
					Enabled bool   `json:"enabled"`
					Value   string `json:"value"`
				} `json:"nameForThisLocation"`
				StateOrProvince struct {
					Value string `json:"value"`
				} `json:"stateOrProvince"`
				TimeZone struct {
					Enabled bool   `json:"enabled"`
					Value   string `json:"value"`
				} `json:"timeZone"`
			} `json:"options"`
		} `json:"companySettings"`
		ContactInfo struct {
			Description string `json:"description"`
			Options     struct {
				Contact struct {
					Options struct {
						Company struct {
							Value string `json:"value"`
						} `json:"company"`
						FirstName struct {
							Value string `json:"value"`
						} `json:"firstName"`
						LastName struct {
							Value string `json:"value"`
						} `json:"lastName"`
						Title struct {
							Value string `json:"value"`
						} `json:"title"`
					} `json:"options"`
				} `json:"contact"`
				ContactNumbers struct {
					Options struct {
						Emergency struct {
							Value string `json:"value"`
						} `json:"emergency"`
						Mobile struct {
							Value string `json:"value"`
						} `json:"mobile"`
						Office struct {
							Value string `json:"value"`
						} `json:"office"`
						PrimaryEmail struct {
							Value string `json:"value"`
						} `json:"primaryEmail"`
						SecondaryEmail struct {
							Value string `json:"value"`
						} `json:"secondaryEmail"`
					} `json:"options"`
				} `json:"contactNumbers"`
			} `json:"options"`
		} `json:"contactInfo"`
		DeviceProvisioning struct {
			Description string `json:"description"`
			Options     struct {
				PhoneProvisioningCertificates struct {
					Port struct {
						Value int `json:"value"`
					} `json:"port"`
					ServerCertificate struct {
						Value int `json:"value"`
					} `json:"serverCertificate"`
				} `json:"phoneProvisioningCertificates"`
			} `json:"options"`
		} `json:"deviceProvisioning"`
		ExternalBilling struct {
			Description string `json:"description"`
			Options     struct {
				Enabled struct {
					Value bool `json:"value"`
				} `json:"enabled"`
				Host struct {
					Description string `json:"description"`
					Value       string `json:"value"`
				} `json:"host"`
				Port struct {
					Description string `json:"description"`
					Value       int    `json:"value"`
				} `json:"port"`
			} `json:"options"`
		} `json:"externalBilling"`
		ExternalMessaging struct {
			Description string `json:"description"`
			Options     struct {
				Mms struct {
					Hosts   []interface{} `json:"hosts"`
					OwnFqdn string        `json:"ownFqdn"`
					Tls     bool          `json:"tls"`
				} `json:"mms"`
				Smpp struct {
					Did []struct {
						SmsDid string `json:"smsDid"`
					} `json:"did"`
					Host     string `json:"host"`
					Password string `json:"password"`
					Port     int    `json:"port"`
					Ssl      bool   `json:"ssl"`
					User     string `json:"user"`
				} `json:"smpp"`
				SmsGate string `json:"smsGate"`
				Zgate   struct {
					AccessToken string `json:"accessToken"`
					Hosts       []struct {
						Host string `json:"host"`
						Port int    `json:"port"`
					} `json:"hosts"`
					Secret string `json:"secret"`
					User   string `json:"user"`
				} `json:"zgate"`
			} `json:"options"`
		} `json:"externalMessaging"`
		LdapConfigurationSettings struct {
			Description string `json:"description"`
			Options     struct {
				Credentials struct {
					Options struct {
						Name struct {
							Value string `json:"value"`
						} `json:"name"`
						Password struct {
							Value string `json:"value"`
						} `json:"password"`
					} `json:"options"`
				} `json:"credentials"`
				Domain struct {
					Value string `json:"value"`
				} `json:"domain"`
				Enabled struct {
					Value bool `json:"value"`
				} `json:"enabled"`
				LdapServers []interface{} `json:"ldapServers"`
				SearchBase  struct {
					Value string `json:"value"`
				} `json:"searchBase"`
				SearchFilter struct {
					Value string `json:"value"`
				} `json:"searchFilter"`
			} `json:"options"`
		} `json:"ldapConfigurationSettings"`
		Miscelaneous struct {
			Description string `json:"description"`
			Options     struct {
				CallRecording struct {
					Description string `json:"description"`
					Options     struct {
						ToneAtStart struct {
							Value bool `json:"value"`
						} `json:"toneAtStart"`
						ToneEveryTime struct {
							Value bool `json:"value"`
						} `json:"toneEveryTime"`
						ToneRepeatPeriod struct {
							Value int `json:"value"`
						} `json:"toneRepeatPeriod"`
					} `json:"options"`
				} `json:"callRecording"`
				FaxSettings struct {
					Description string `json:"description"`
					Options     struct {
						CompanyFax struct {
							Value string `json:"value"`
						} `json:"companyFax"`
						CompanyName struct {
							Value string `json:"value"`
						} `json:"companyName"`
						FaxFileFormat          string `json:"faxFileFormat"`
						IntervalBetweenRetries struct {
							Value int `json:"value"`
						} `json:"intervalBetweenRetries"`
						NumberOfRetries struct {
							Value int `json:"value"`
						} `json:"numberOfRetries"`
					} `json:"options"`
				} `json:"faxSettings"`
				ForwardCLIpolicy struct {
					Value string `json:"value"`
				} `json:"forwardCLIpolicy"`
				ShutdownOnPowerLoss struct {
					Value   bool `json:"value"`
					Visible bool `json:"visible"`
				} `json:"shutdownOnPowerLoss"`
			} `json:"options"`
		} `json:"miscelaneous"`
		NetworkSettings struct {
			Description string `json:"description"`
			Options     struct {
				DhcpCfg struct {
					Enabled bool `json:"enabled"`
					Options struct {
						IsUsed struct {
							Value bool `json:"value"`
						} `json:"isUsed"`
						Mask struct {
							Value string `json:"value"`
						} `json:"mask"`
						Opt struct {
							DefDomain struct {
								Value string `json:"value"`
							} `json:"defDomain"`
							DefGateway struct {
								Value string `json:"value"`
							} `json:"defGateway"`
							DefLeaseTime struct {
								Value int `json:"value"`
							} `json:"defLeaseTime"`
							Dns struct {
								Value string `json:"value"`
							} `json:"dns"`
							DnsDefault struct {
								Value bool `json:"value"`
							} `json:"dnsDefault"`
							Ntp struct {
								Value string `json:"value"`
							} `json:"ntp"`
							NtpDefault struct {
								Value bool `json:"value"`
							} `json:"ntpDefault"`
							TimeOffset struct {
								Value int `json:"value"`
							} `json:"timeOffset"`
						} `json:"opt"`
						Range struct {
							First struct {
								Value string `json:"value"`
							} `json:"first"`
							Last struct {
								Value string `json:"value"`
							} `json:"last"`
						} `json:"range"`
						Subnet struct {
							Value string `json:"value"`
						} `json:"subnet"`
					} `json:"options"`
				} `json:"dhcpCfg"`
				DnsServers struct {
					Description string `json:"description"`
					Options     struct {
						DnsServ1 struct {
							Value string `json:"value"`
						} `json:"dnsServ1"`
						DnsServ2 struct {
							Value string `json:"value"`
						} `json:"dnsServ2"`
					} `json:"options"`
				} `json:"dnsServers"`
				IpAddress struct {
					Description string `json:"description"`
					Enabled     bool   `json:"enabled"`
					Options     struct {
						Config struct {
							Description string `json:"description"`
							Enabled     bool   `json:"enabled"`
							Value       string `json:"value"`
						} `json:"config"`
						Gw struct {
							Value string `json:"value"`
						} `json:"gw"`
						Ip struct {
							Value string `json:"value"`
						} `json:"ip"`
						Mask struct {
							Value string `json:"value"`
						} `json:"mask"`
					} `json:"options"`
				} `json:"ipAddress"`
			} `json:"options"`
		} `json:"networkSettings"`
		ProvisioningServer struct {
			Description string `json:"description"`
			Options     struct {
				ConfigSubdir struct {
					Description string `json:"description"`
					Value       string `json:"value"`
				} `json:"configSubdir"`
				ExternalFtp struct {
					Description string `json:"description"`
					Options     struct {
						Address struct {
							Value string `json:"value"`
						} `json:"address"`
						Dir struct {
							Value string `json:"value"`
						} `json:"dir"`
						Login struct {
							Value string `json:"value"`
						} `json:"login"`
						Password struct {
							Value string `json:"value"`
						} `json:"password"`
					} `json:"options"`
				} `json:"externalFtp"`
				InternalServer struct {
					Value bool `json:"value"`
				} `json:"internalServer"`
			} `json:"options"`
		} `json:"provisioningServer"`
		ProxiesSettings struct {
			Description string `json:"description"`
			Options     struct {
				ProxyUI struct {
					Description string `json:"description"`
					Options     struct {
						Address struct {
							Value string `json:"value"`
						} `json:"address"`
						AuthenticationRequired struct {
							Value bool `json:"value"`
						} `json:"authenticationRequired"`
						Password struct {
							Value string `json:"value"`
						} `json:"password"`
						Port struct {
							Value int `json:"value"`
						} `json:"port"`
						Protocol struct {
							Value string `json:"value"`
						} `json:"protocol"`
						UseProxy struct {
							Value string `json:"value"`
						} `json:"useProxy"`
						UserName struct {
							Value string `json:"value"`
						} `json:"userName"`
					} `json:"options"`
				} `json:"proxyUI"`
			} `json:"options"`
		} `json:"proxiesSettings"`
		SecuritySettings struct {
			Description string `json:"description"`
			Options     struct {
				Passwords struct {
					Description string `json:"description"`
					Options     struct {
						BackwardCompatibilityMode struct {
							Value bool `json:"value"`
						} `json:"backwardCompatibilityMode"`
						DataRetentionPolicy struct {
							Description string `json:"description"`
							Options     struct {
								RetentionPeriod struct {
									Value int `json:"value"`
								} `json:"retentionPeriod"`
							} `json:"options"`
						} `json:"dataRetentionPolicy"`
						Mfa struct {
							Description string `json:"description"`
							Options     struct {
								Administrators struct {
									Description string `json:"description"`
									Options     struct {
										MfaMode struct {
											Value bool `json:"value"`
										} `json:"mfaMode"`
										Transport struct {
											Value int `json:"value"`
										} `json:"transport"`
										TrustedDeviceMode struct {
											Value bool `json:"value"`
										} `json:"trustedDeviceMode"`
									} `json:"options"`
								} `json:"administrators"`
								OneTimeCode struct {
									Description string `json:"description"`
									Options     struct {
										Length struct {
											Value int `json:"value"`
										} `json:"length"`
										LifeTime struct {
											Value int `json:"value"`
										} `json:"lifeTime"`
										MessageTemplate struct {
											Value string `json:"value"`
										} `json:"messageTemplate"`
										SubjectTemplate struct {
											Value string `json:"value"`
										} `json:"subjectTemplate"`
										Tries struct {
											Value int `json:"value"`
										} `json:"tries"`
									} `json:"options"`
								} `json:"oneTimeCode"`
								TrustedDevice struct {
									Description string `json:"description"`
									Options     struct {
										LifeTime struct {
											Value int `json:"value"`
										} `json:"lifeTime"`
									} `json:"options"`
								} `json:"trustedDevice"`
								Users struct {
									Description string `json:"description"`
									Options     struct {
										MfaMode struct {
											Value bool `json:"value"`
										} `json:"mfaMode"`
									} `json:"options"`
								} `json:"users"`
							} `json:"options"`
						} `json:"mfa"`
						Password struct {
							Description string `json:"description"`
							Options     struct {
								DefaultPassword struct {
									Value string `json:"value"`
								} `json:"defaultPassword"`
								DisableDefaultPassword struct {
									Value bool `json:"value"`
								} `json:"disableDefaultPassword"`
								ExpirationPeriod struct {
									Value int `json:"value"`
								} `json:"expirationPeriod"`
								MinimumLength struct {
									Value int `json:"value"`
								} `json:"minimumLength"`
								PrlLifetime struct {
									Value int `json:"value"`
								} `json:"prlLifetime"`
								PrlMessageTemplate struct {
									Value string `json:"value"`
								} `json:"prlMessageTemplate"`
								PrlMode struct {
									Value bool `json:"value"`
								} `json:"prlMode"`
								PrlSubjectTemplate struct {
									Value string `json:"value"`
								} `json:"prlSubjectTemplate"`
							} `json:"options"`
						} `json:"password"`
						Pin struct {
							Description string `json:"description"`
							Options     struct {
								DefaultPassword struct {
									Value string `json:"value"`
								} `json:"defaultPassword"`
								MinimumLength struct {
									Value int `json:"value"`
								} `json:"minimumLength"`
							} `json:"options"`
						} `json:"pin"`
					} `json:"options"`
				} `json:"passwords"`
				WebServer struct {
					Description string `json:"description"`
					Options     struct {
						TrustedOrigins struct {
							Value []interface{} `json:"value"`
						} `json:"trustedOrigins"`
					} `json:"options"`
				} `json:"webServer"`
			} `json:"options"`
		} `json:"securitySettings"`
		Smtp struct {
			Description string `json:"description"`
			Options     struct {
				CodePage struct {
					Description string `json:"description"`
					Enabled     bool   `json:"enabled"`
					Value       string `json:"value"`
					Visible     bool   `json:"visible"`
				} `json:"codePage"`
				Enabled struct {
					Value bool `json:"value"`
				} `json:"enabled"`
				MaxAttachSize struct {
					Value int `json:"value"`
				} `json:"maxAttachSize"`
				Password struct {
					Value string `json:"value"`
				} `json:"password"`
				Port struct {
					Value int `json:"value"`
				} `json:"port"`
				Protocol struct {
					Value string `json:"value"`
				} `json:"protocol"`
				RequireAuth struct {
					Value bool `json:"value"`
				} `json:"requireAuth"`
				SmtpServer struct {
					Value string `json:"value"`
				} `json:"smtpServer"`
				SystemEmailAddress struct {
					Value string `json:"value"`
				} `json:"systemEmailAddress"`
				UserName struct {
					Value string `json:"value"`
				} `json:"userName"`
			} `json:"options"`
		} `json:"smtp"`
		Tts struct {
			Description string `json:"description"`
			Options     struct {
				Password string        `json:"password"`
				Servers  []interface{} `json:"servers"`
				UserName string        `json:"userName"`
			} `json:"options"`
		} `json:"tts"`
	} `json:"systemSettings"`
}

func HandleSystemSettingsResponse(resp *http.Response) (*SystemSettings, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var licenseResponse SystemSettings
	err = json.Unmarshal(body, &licenseResponse)
	if err != nil {
		return nil, err
	}

	return &licenseResponse, nil
}
