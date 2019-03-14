package basic

type Service struct {
	Accessdown                   string `json:"accessdown,omitempty"`
	All                          bool   `json:"all,omitempty"`
	Appflowlog                   string `json:"appflowlog,omitempty"`
	Cacheable                    string `json:"cacheable,omitempty"`
	Cachetype                    string `json:"cachetype,omitempty"`
	Cip                          string `json:"cip,omitempty"`
	Cipheader                    string `json:"cipheader,omitempty"`
	Cka                          string `json:"cka,omitempty"`
	Cleartextport                int    `json:"cleartextport,omitempty"`
	Clmonowner                   int    `json:"clmonowner,omitempty"`
	Clmonview                    int    `json:"clmonview,omitempty"`
	Clttimeout                   int    `json:"clttimeout,omitempty"`
	Cmp                          string `json:"cmp,omitempty"`
	Comment                      string `json:"comment,omitempty"`
	Contentinspectionprofilename string `json:"contentinspectionprofilename,omitempty"`
	Customserverid               string `json:"customserverid,omitempty"`
	Delay                        int    `json:"delay,omitempty"`
	Dnsprofilename               string `json:"dnsprofilename,omitempty"`
	Downstateflush               string `json:"downstateflush,omitempty"`
	Dupstate                     string `json:"dup_state,omitempty"`
	Graceful                     string `json:"graceful,omitempty"`
	Gslb                         string `json:"gslb,omitempty"`
	Hashid                       int    `json:"hashid,omitempty"`
	Healthmonitor                string `json:"healthmonitor,omitempty"`
	Httpprofilename              string `json:"httpprofilename,omitempty"`
	Internal                     bool   `json:"Internal,omitempty"`
	Ip                           string `json:"ip,omitempty"`
	Ipaddress                    string `json:"ipaddress,omitempty"`
	Lastresponse                 string `json:"lastresponse,omitempty"`
	Maxbandwidth                 int    `json:"maxbandwidth,omitempty"`
	Maxclient                    int    `json:"maxclient,omitempty"`
	Maxreq                       int    `json:"maxreq,omitempty"`
	Monconnectionclose           string `json:"monconnectionclose,omitempty"`
	Monitornamesvc               string `json:"monitor_name_svc,omitempty"`
	Monitorstate                 string `json:"monitor_state,omitempty"`
	Monstatcode                  int    `json:"monstatcode,omitempty"`
	Monstatparam1                int    `json:"monstatparam1,omitempty"`
	Monstatparam2                int    `json:"monstatparam2,omitempty"`
	Monstatparam3                int    `json:"monstatparam3,omitempty"`
	Monthreshold                 int    `json:"monthreshold,omitempty"`
	Name                         string `json:"name,omitempty"`
	Netprofile                   string `json:"netprofile,omitempty"`
	Newname                      string `json:"newname,omitempty"`
	Nodefaultbindings            string `json:"nodefaultbindings,omitempty"`
	Numofconnections             int    `json:"numofconnections,omitempty"`
	Oracleserverversion          string `json:"oracleserverversion,omitempty"`
	Pathmonitor                  string `json:"pathmonitor,omitempty"`
	Pathmonitorindv              string `json:"pathmonitorindv,omitempty"`
	Policyname                   string `json:"policyname,omitempty"`
	Port                         int    `json:"port,omitempty"`
	Processlocal                 string `json:"processlocal,omitempty"`
	Publicip                     string `json:"publicip,omitempty"`
	Publicport                   int    `json:"publicport,omitempty"`
	Responsetime                 int    `json:"responsetime,omitempty"`
	Riseapbrstatsmsgcode         int    `json:"riseapbrstatsmsgcode,omitempty"`
	Riseapbrstatsmsgcode2        int    `json:"riseapbrstatsmsgcode2,omitempty"`
	Rtspsessionidremap           string `json:"rtspsessionidremap,omitempty"`
	Sc                           string `json:"sc,omitempty"`
	Serverid                     int    `json:"serverid,omitempty"`
	Servername                   string `json:"servername,omitempty"`
	Serviceconftype              bool   `json:"serviceconftype,omitempty"`
	Serviceconftype2             string `json:"serviceconftype2,omitempty"`
	Serviceipstr                 string `json:"serviceipstr,omitempty"`
	Servicetype                  string `json:"servicetype,omitempty"`
	Sp                           string `json:"sp,omitempty"`
	State                        string `json:"state,omitempty"`
	Statechangetimemsec          int    `json:"statechangetimemsec,omitempty"`
	Statechangetimesec           string `json:"statechangetimesec,omitempty"`
	Stateupdatereason            int    `json:"stateupdatereason,omitempty"`
	Svrstate                     string `json:"svrstate,omitempty"`
	Svrtimeout                   int    `json:"svrtimeout,omitempty"`
	Tcpb                         string `json:"tcpb,omitempty"`
	Tcpprofilename               string `json:"tcpprofilename,omitempty"`
	Td                           int    `json:"td,omitempty"`
	Tickssincelaststatechange    int    `json:"tickssincelaststatechange,omitempty"`
	Useproxyport                 string `json:"useproxyport,omitempty"`
	Usip                         string `json:"usip,omitempty"`
	Value                        string `json:"value,omitempty"`
	Weight                       int    `json:"weight,omitempty"`
}
