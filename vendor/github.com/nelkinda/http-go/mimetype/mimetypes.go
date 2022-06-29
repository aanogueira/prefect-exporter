// Package mimetype defines the constants for all IANA registered MIME Types.
// See https://www.iana.org/assignments/media-types for the official list.
// The advantage of using the constants from this package over string literals is correctness.
// The compiler can spot errors that the runtime couldn't.
// Example:
//     // Spelling mistakes, not caught by the compiler
//     w.Header().Add("Content-Tpye", "applicaiton/xml")
//
//     // Spelling mistakes, caught by the compiler
//     w.Header().Add(header.ContentTpye, mimetype.ApplicatoinXml)
//
// Note: Currently, only a subset is implemented.
// The goal is to implement all registered MIME Types by scraping.
package mimetype

const (
	// Application1dInterleavedParityfec is the constant for the mime type "application/1d-interleaved-parityfec".
	Application1dInterleavedParityfec = "application/1d-interleaved-parityfec"

	// Application3gpdashQoeReportXml is the constant for the mime type "application/3gpdash-qoe-report+xml".
	Application3gpdashQoeReportXml = "application/3gpdash-qoe-report+xml"

	// Application3gppImsXml is the constant for the mime type "application/3gpp-ims+xml".
	Application3gppImsXml = "application/3gpp-ims+xml"

	// ApplicationA2L is the constant for the mime type "application/A2L".
	ApplicationA2L = "application/A2L"

	// ApplicationActivemessage is the constant for the mime type "application/activemessage".
	ApplicationActivemessage = "application/activemessage"

	// ApplicationActivityJson is the constant for the mime type "application/activity+json".
	ApplicationActivityJson = "application/activity+json"

	// ApplicationAltoCostmapJson is the constant for the mime type "application/alto-costmap+json".
	ApplicationAltoCostmapJson = "application/alto-costmap+json"

	// ApplicationAltoCostmapfilterJson is the constant for the mime type "application/alto-costmapfilter+json".
	ApplicationAltoCostmapfilterJson = "application/alto-costmapfilter+json"

	// ApplicationAltoDirectoryJson is the constant for the mime type "application/alto-directory+json".
	ApplicationAltoDirectoryJson = "application/alto-directory+json"

	// ApplicationAltoEndpointpropJson is the constant for the mime type "application/alto-endpointprop+json".
	ApplicationAltoEndpointpropJson = "application/alto-endpointprop+json"

	// ApplicationAltoEndpointpropparamsJson is the constant for the mime type "application/alto-endpointpropparams+json".
	ApplicationAltoEndpointpropparamsJson = "application/alto-endpointpropparams+json"

	// ApplicationAltoEndpointcostJson is the constant for the mime type "application/alto-endpointcost+json".
	ApplicationAltoEndpointcostJson = "application/alto-endpointcost+json"

	// ApplicationAltoEndpointcostparamsJson is the constant for the mime type "application/alto-endpointcostparams+json".
	ApplicationAltoEndpointcostparamsJson = "application/alto-endpointcostparams+json"

	// ApplicationAltoErrorJson is the constant for the mime type "application/alto-error+json".
	ApplicationAltoErrorJson = "application/alto-error+json"

	// ApplicationAltoNetworkmapfilterJson is the constant for the mime type "application/alto-networkmapfilter+json".
	ApplicationAltoNetworkmapfilterJson = "application/alto-networkmapfilter+json"

	// ApplicationAltoNetworkmapJson is the constant for the mime type "application/alto-networkmap+json".
	ApplicationAltoNetworkmapJson = "application/alto-networkmap+json"

	// ApplicationAML is the constant for the mime type "application/AML".
	ApplicationAML = "application/AML"

	// ApplicationAndrewInset is the constant for the mime type "application/andrew-inset".
	ApplicationAndrewInset = "application/andrew-inset"

	// ApplicationApplefile is the constant for the mime type "application/applefile".
	ApplicationApplefile = "application/applefile"

	// ApplicationATF is the constant for the mime type "application/ATF".
	ApplicationATF = "application/ATF"

	// ApplicationATFX is the constant for the mime type "application/ATFX".
	ApplicationATFX = "application/ATFX"

	// ApplicationAtomXml is the constant for the mime type "application/atom+xml".
	ApplicationAtomXml = "application/atom+xml"

	// ApplicationAtomcatXml is the constant for the mime type "application/atomcat+xml".
	ApplicationAtomcatXml = "application/atomcat+xml"

	// ApplicationAtomdeletedXml is the constant for the mime type "application/atomdeleted+xml".
	ApplicationAtomdeletedXml = "application/atomdeleted+xml"

	// ApplicationAtomicmail is the constant for the mime type "application/atomicmail".
	ApplicationAtomicmail = "application/atomicmail"

	// ApplicationAtomsvcXml is the constant for the mime type "application/atomsvc+xml".
	ApplicationAtomsvcXml = "application/atomsvc+xml"

	// ApplicationAtscDwdXml is the constant for the mime type "application/atsc-dwd+xml".
	ApplicationAtscDwdXml = "application/atsc-dwd+xml"

	// ApplicationAtscHeldXml is the constant for the mime type "application/atsc-held+xml".
	ApplicationAtscHeldXml = "application/atsc-held+xml"

	// ApplicationAtscRsatXml is the constant for the mime type "application/atsc-rsat+xml".
	ApplicationAtscRsatXml = "application/atsc-rsat+xml"

	// ApplicationATXML is the constant for the mime type "application/ATXML".
	ApplicationATXML = "application/ATXML"

	// ApplicationAuthPolicyXml is the constant for the mime type "application/auth-policy+xml".
	ApplicationAuthPolicyXml = "application/auth-policy+xml"

	// ApplicationBacnetXddZip is the constant for the mime type "application/bacnet-xdd+zip".
	ApplicationBacnetXddZip = "application/bacnet-xdd+zip"

	// ApplicationBatchSMTP is the constant for the mime type "application/batch-SMTP".
	ApplicationBatchSMTP = "application/batch-SMTP"

	// ApplicationBeepXml is the constant for the mime type "application/beep+xml".
	ApplicationBeepXml = "application/beep+xml"

	// ApplicationCalendarJson is the constant for the mime type "application/calendar+json".
	ApplicationCalendarJson = "application/calendar+json"

	// ApplicationCalendarXml is the constant for the mime type "application/calendar+xml".
	ApplicationCalendarXml = "application/calendar+xml"

	// ApplicationCallCompletion is the constant for the mime type "application/call-completion".
	ApplicationCallCompletion = "application/call-completion"

	// ApplicationCALS1840 is the constant for the mime type "application/CALS-1840".
	ApplicationCALS1840 = "application/CALS-1840"

	// ApplicationCbor is the constant for the mime type "application/cbor".
	ApplicationCbor = "application/cbor"

	// ApplicationCccex is the constant for the mime type "application/cccex".
	ApplicationCccex = "application/cccex"

	// ApplicationCcmpXml is the constant for the mime type "application/ccmp+xml".
	ApplicationCcmpXml = "application/ccmp+xml"

	// ApplicationCcxmlXml is the constant for the mime type "application/ccxml+xml".
	ApplicationCcxmlXml = "application/ccxml+xml"

	// ApplicationCDFXXML is the constant for the mime type "application/CDFX+XML".
	ApplicationCDFXXML = "application/CDFX+XML"

	// ApplicationCdmiCapability is the constant for the mime type "application/cdmi-capability".
	ApplicationCdmiCapability = "application/cdmi-capability"

	// ApplicationCdmiContainer is the constant for the mime type "application/cdmi-container".
	ApplicationCdmiContainer = "application/cdmi-container"

	// ApplicationCdmiDomain is the constant for the mime type "application/cdmi-domain".
	ApplicationCdmiDomain = "application/cdmi-domain"

	// ApplicationCdmiObject is the constant for the mime type "application/cdmi-object".
	ApplicationCdmiObject = "application/cdmi-object"

	// ApplicationCdmiQueue is the constant for the mime type "application/cdmi-queue".
	ApplicationCdmiQueue = "application/cdmi-queue"

	// ApplicationCdni is the constant for the mime type "application/cdni".
	ApplicationCdni = "application/cdni"

	// ApplicationCEA is the constant for the mime type "application/CEA".
	ApplicationCEA = "application/CEA"

	// ApplicationCea2018Xml is the constant for the mime type "application/cea-2018+xml".
	ApplicationCea2018Xml = "application/cea-2018+xml"

	// ApplicationCellmlXml is the constant for the mime type "application/cellml+xml".
	ApplicationCellmlXml = "application/cellml+xml"

	// ApplicationCfw is the constant for the mime type "application/cfw".
	ApplicationCfw = "application/cfw"

	// ApplicationClueInfoXml is the constant for the mime type "application/clue_info+xml".
	ApplicationClueInfoXml = "application/clue_info+xml"

	// ApplicationCms is the constant for the mime type "application/cms".
	ApplicationCms = "application/cms"

	// ApplicationCnrpXml is the constant for the mime type "application/cnrp+xml".
	ApplicationCnrpXml = "application/cnrp+xml"

	// ApplicationCoapGroupJson is the constant for the mime type "application/coap-group+json".
	ApplicationCoapGroupJson = "application/coap-group+json"

	// ApplicationCoapPayload is the constant for the mime type "application/coap-payload".
	ApplicationCoapPayload = "application/coap-payload"

	// ApplicationCommonground is the constant for the mime type "application/commonground".
	ApplicationCommonground = "application/commonground"

	// ApplicationConferenceInfoXml is the constant for the mime type "application/conference-info+xml".
	ApplicationConferenceInfoXml = "application/conference-info+xml"

	// ApplicationCplXml is the constant for the mime type "application/cpl+xml".
	ApplicationCplXml = "application/cpl+xml"

	// ApplicationCose is the constant for the mime type "application/cose".
	ApplicationCose = "application/cose"

	// ApplicationCoseKey is the constant for the mime type "application/cose-key".
	ApplicationCoseKey = "application/cose-key"

	// ApplicationCoseKeySet is the constant for the mime type "application/cose-key-set".
	ApplicationCoseKeySet = "application/cose-key-set"

	// ApplicationCsrattrs is the constant for the mime type "application/csrattrs".
	ApplicationCsrattrs = "application/csrattrs"

	// ApplicationCstaXml is the constant for the mime type "application/csta+xml".
	ApplicationCstaXml = "application/csta+xml"

	// ApplicationCSTAdataXml is the constant for the mime type "application/CSTAdata+xml".
	ApplicationCSTAdataXml = "application/CSTAdata+xml"

	// ApplicationCsvmJson is the constant for the mime type "application/csvm+json".
	ApplicationCsvmJson = "application/csvm+json"

	// ApplicationCwt is the constant for the mime type "application/cwt".
	ApplicationCwt = "application/cwt"

	// ApplicationCybercash is the constant for the mime type "application/cybercash".
	ApplicationCybercash = "application/cybercash"

	// ApplicationDashXml is the constant for the mime type "application/dash+xml".
	ApplicationDashXml = "application/dash+xml"

	// ApplicationDashdelta is the constant for the mime type "application/dashdelta".
	ApplicationDashdelta = "application/dashdelta"

	// ApplicationDavmountXml is the constant for the mime type "application/davmount+xml".
	ApplicationDavmountXml = "application/davmount+xml"

	// ApplicationDcaRft is the constant for the mime type "application/dca-rft".
	ApplicationDcaRft = "application/dca-rft"

	// ApplicationDCD is the constant for the mime type "application/DCD".
	ApplicationDCD = "application/DCD"

	// ApplicationDecDx is the constant for the mime type "application/dec-dx".
	ApplicationDecDx = "application/dec-dx"

	// ApplicationDialogInfoXml is the constant for the mime type "application/dialog-info+xml".
	ApplicationDialogInfoXml = "application/dialog-info+xml"

	// ApplicationDicom is the constant for the mime type "application/dicom".
	ApplicationDicom = "application/dicom"

	// ApplicationDicomJson is the constant for the mime type "application/dicom+json".
	ApplicationDicomJson = "application/dicom+json"

	// ApplicationDicomXml is the constant for the mime type "application/dicom+xml".
	ApplicationDicomXml = "application/dicom+xml"

	// ApplicationDII is the constant for the mime type "application/DII".
	ApplicationDII = "application/DII"

	// ApplicationDIT is the constant for the mime type "application/DIT".
	ApplicationDIT = "application/DIT"

	// ApplicationDns is the constant for the mime type "application/dns".
	ApplicationDns = "application/dns"

	// ApplicationDnsJson is the constant for the mime type "application/dns+json".
	ApplicationDnsJson = "application/dns+json"

	// ApplicationDnsMessage is the constant for the mime type "application/dns-message".
	ApplicationDnsMessage = "application/dns-message"

	// ApplicationDskppXml is the constant for the mime type "application/dskpp+xml".
	ApplicationDskppXml = "application/dskpp+xml"

	// ApplicationDsscDer is the constant for the mime type "application/dssc+der".
	ApplicationDsscDer = "application/dssc+der"

	// ApplicationDsscXml is the constant for the mime type "application/dssc+xml".
	ApplicationDsscXml = "application/dssc+xml"

	// ApplicationDvcs is the constant for the mime type "application/dvcs".
	ApplicationDvcs = "application/dvcs"

	// ApplicationEcmascript is the constant for the mime type "application/ecmascript".
	ApplicationEcmascript = "application/ecmascript"

	// ApplicationEDIConsent is the constant for the mime type "application/EDI-consent".
	ApplicationEDIConsent = "application/EDI-consent"

	// ApplicationEDIFACT is the constant for the mime type "application/EDIFACT".
	ApplicationEDIFACT = "application/EDIFACT"

	// ApplicationEDIX12 is the constant for the mime type "application/EDI-X12".
	ApplicationEDIX12 = "application/EDI-X12"

	// ApplicationEfi is the constant for the mime type "application/efi".
	ApplicationEfi = "application/efi"

	// ApplicationEmergencyCallDataCommentXml is the constant for the mime type "application/EmergencyCallData.Comment+xml".
	ApplicationEmergencyCallDataCommentXml = "application/EmergencyCallData.Comment+xml"

	// ApplicationEmergencyCallDataControlXml is the constant for the mime type "application/EmergencyCallData.Control+xml".
	ApplicationEmergencyCallDataControlXml = "application/EmergencyCallData.Control+xml"

	// ApplicationEmergencyCallDataDeviceInfoXml is the constant for the mime type "application/EmergencyCallData.DeviceInfo+xml".
	ApplicationEmergencyCallDataDeviceInfoXml = "application/EmergencyCallData.DeviceInfo+xml"

	// ApplicationEmergencyCallDataECallMSD is the constant for the mime type "application/EmergencyCallData.eCall.MSD".
	ApplicationEmergencyCallDataECallMSD = "application/EmergencyCallData.eCall.MSD"

	// ApplicationEmergencyCallDataProviderInfoXml is the constant for the mime type "application/EmergencyCallData.ProviderInfo+xml".
	ApplicationEmergencyCallDataProviderInfoXml = "application/EmergencyCallData.ProviderInfo+xml"

	// ApplicationEmergencyCallDataServiceInfoXml is the constant for the mime type "application/EmergencyCallData.ServiceInfo+xml".
	ApplicationEmergencyCallDataServiceInfoXml = "application/EmergencyCallData.ServiceInfo+xml"

	// ApplicationEmergencyCallDataSubscriberInfoXml is the constant for the mime type "application/EmergencyCallData.SubscriberInfo+xml".
	ApplicationEmergencyCallDataSubscriberInfoXml = "application/EmergencyCallData.SubscriberInfo+xml"

	// ApplicationEmergencyCallDataVEDSXml is the constant for the mime type "application/EmergencyCallData.VEDS+xml".
	ApplicationEmergencyCallDataVEDSXml = "application/EmergencyCallData.VEDS+xml"

	// ApplicationEmmaXml is the constant for the mime type "application/emma+xml".
	ApplicationEmmaXml = "application/emma+xml"

	// ApplicationEmotionmlXml is the constant for the mime type "application/emotionml+xml".
	ApplicationEmotionmlXml = "application/emotionml+xml"

	// ApplicationEncaprtp is the constant for the mime type "application/encaprtp".
	ApplicationEncaprtp = "application/encaprtp"

	// ApplicationEppXml is the constant for the mime type "application/epp+xml".
	ApplicationEppXml = "application/epp+xml"

	// ApplicationEpubZip is the constant for the mime type "application/epub+zip".
	ApplicationEpubZip = "application/epub+zip"

	// ApplicationEshop is the constant for the mime type "application/eshop".
	ApplicationEshop = "application/eshop"

	// ApplicationExample is the constant for the mime type "application/example".
	ApplicationExample = "application/example"

	// ApplicationExi is the constant for the mime type "application/exi".
	ApplicationExi = "application/exi"

	// ApplicationExpectCtReportJson is the constant for the mime type "application/expect-ct-report+json".
	ApplicationExpectCtReportJson = "application/expect-ct-report+json"

	// ApplicationFastinfoset is the constant for the mime type "application/fastinfoset".
	ApplicationFastinfoset = "application/fastinfoset"

	// ApplicationFastsoap is the constant for the mime type "application/fastsoap".
	ApplicationFastsoap = "application/fastsoap"

	// ApplicationFdtXml is the constant for the mime type "application/fdt+xml".
	ApplicationFdtXml = "application/fdt+xml"

	// ApplicationFhirJson is the constant for the mime type "application/fhir+json".
	ApplicationFhirJson = "application/fhir+json"

	// ApplicationFhirXml is the constant for the mime type "application/fhir+xml".
	ApplicationFhirXml = "application/fhir+xml"

	// ApplicationFits is the constant for the mime type "application/fits".
	ApplicationFits = "application/fits"

	// ApplicationFlexfec is the constant for the mime type "application/flexfec".
	ApplicationFlexfec = "application/flexfec"

	// ApplicationFontSfnt is the constant for the mime type "application/font-sfnt - DEPRECATED in favor of font/sfnt".
	ApplicationFontSfnt = "application/font-sfnt"

	// ApplicationFontTdpfr is the constant for the mime type "application/font-tdpfr".
	ApplicationFontTdpfr = "application/font-tdpfr"

	// ApplicationFontWoff is the constant for the mime type "application/font-woff - DEPRECATED in favor of font/woff".
	ApplicationFontWoff = "application/font-woff"

	// ApplicationFrameworkAttributesXml is the constant for the mime type "application/framework-attributes+xml".
	ApplicationFrameworkAttributesXml = "application/framework-attributes+xml"

	// ApplicationGeoJson is the constant for the mime type "application/geo+json".
	ApplicationGeoJson = "application/geo+json"

	// ApplicationGeoJsonSeq is the constant for the mime type "application/geo+json-seq".
	ApplicationGeoJsonSeq = "application/geo+json-seq"

	// ApplicationGeopackageSqlite3 is the constant for the mime type "application/geopackage+sqlite3".
	ApplicationGeopackageSqlite3 = "application/geopackage+sqlite3"

	// ApplicationGeoxacmlXml is the constant for the mime type "application/geoxacml+xml".
	ApplicationGeoxacmlXml = "application/geoxacml+xml"

	// ApplicationGltfBuffer is the constant for the mime type "application/gltf-buffer".
	ApplicationGltfBuffer = "application/gltf-buffer"

	// ApplicationGmlXml is the constant for the mime type "application/gml+xml".
	ApplicationGmlXml = "application/gml+xml"

	// ApplicationGzip is the constant for the mime type "application/gzip".
	ApplicationGzip = "application/gzip"

	// ApplicationH224 is the constant for the mime type "application/H224".
	ApplicationH224 = "application/H224"

	// ApplicationHeldXml is the constant for the mime type "application/held+xml".
	ApplicationHeldXml = "application/held+xml"

	// ApplicationHttp is the constant for the mime type "application/http".
	ApplicationHttp = "application/http"

	// ApplicationHyperstudio is the constant for the mime type "application/hyperstudio".
	ApplicationHyperstudio = "application/hyperstudio"

	// ApplicationIbeKeyRequestXml is the constant for the mime type "application/ibe-key-request+xml".
	ApplicationIbeKeyRequestXml = "application/ibe-key-request+xml"

	// ApplicationIbePkgReplyXml is the constant for the mime type "application/ibe-pkg-reply+xml".
	ApplicationIbePkgReplyXml = "application/ibe-pkg-reply+xml"

	// ApplicationIbePpData is the constant for the mime type "application/ibe-pp-data".
	ApplicationIbePpData = "application/ibe-pp-data"

	// ApplicationIges is the constant for the mime type "application/iges".
	ApplicationIges = "application/iges"

	// ApplicationImIscomposingXml is the constant for the mime type "application/im-iscomposing+xml".
	ApplicationImIscomposingXml = "application/im-iscomposing+xml"

	// ApplicationIndex is the constant for the mime type "application/index".
	ApplicationIndex = "application/index"

	// ApplicationIndexCmd is the constant for the mime type "application/index.cmd".
	ApplicationIndexCmd = "application/index.cmd"

	// ApplicationIndexObj is the constant for the mime type "application/index.obj".
	ApplicationIndexObj = "application/index.obj"

	// ApplicationIndexResponse is the constant for the mime type "application/index.response".
	ApplicationIndexResponse = "application/index.response"

	// ApplicationIndexVnd is the constant for the mime type "application/index.vnd".
	ApplicationIndexVnd = "application/index.vnd"

	// ApplicationInkmlXml is the constant for the mime type "application/inkml+xml".
	ApplicationInkmlXml = "application/inkml+xml"

	// ApplicationIOTP is the constant for the mime type "application/IOTP".
	ApplicationIOTP = "application/IOTP"

	// ApplicationIpfix is the constant for the mime type "application/ipfix".
	ApplicationIpfix = "application/ipfix"

	// ApplicationIpp is the constant for the mime type "application/ipp".
	ApplicationIpp = "application/ipp"

	// ApplicationIsup is the constant for the mime type "application/isup".
	ApplicationIsup = "application/isup"

	// ApplicationItsXml is the constant for the mime type "application/its+xml".
	ApplicationItsXml = "application/its+xml"

	// ApplicationJavascript is the constant for the mime type "application/javascript".
	ApplicationJavascript = "application/javascript"

	// ApplicationJf2FeedJson is the constant for the mime type "application/jf2feed+json".
	ApplicationJf2FeedJson = "application/jf2feed+json"

	// ApplicationJose is the constant for the mime type "application/jose".
	ApplicationJose = "application/jose"

	// ApplicationJoseJson is the constant for the mime type "application/jose+json".
	ApplicationJoseJson = "application/jose+json"

	// ApplicationJrdJson is the constant for the mime type "application/jrd+json".
	ApplicationJrdJson = "application/jrd+json"

	// ApplicationJson is the constant for the mime type "application/json".
	ApplicationJson = "application/json"

	// ApplicationJsonPatchJson is the constant for the mime type "application/json-patch+json".
	ApplicationJsonPatchJson = "application/json-patch+json"

	// ApplicationJsonSeq is the constant for the mime type "application/json-seq".
	ApplicationJsonSeq = "application/json-seq"

	// ApplicationJwkJson is the constant for the mime type "application/jwk+json".
	ApplicationJwkJson = "application/jwk+json"

	// ApplicationJwkSetJson is the constant for the mime type "application/jwk-set+json".
	ApplicationJwkSetJson = "application/jwk-set+json"

	// ApplicationJwt is the constant for the mime type "application/jwt".
	ApplicationJwt = "application/jwt"

	// ApplicationKpmlRequestXml is the constant for the mime type "application/kpml-request+xml".
	ApplicationKpmlRequestXml = "application/kpml-request+xml"

	// ApplicationKpmlResponseXml is the constant for the mime type "application/kpml-response+xml".
	ApplicationKpmlResponseXml = "application/kpml-response+xml"

	// ApplicationLdJson is the constant for the mime type "application/ld+json".
	ApplicationLdJson = "application/ld+json"

	// ApplicationLgrXml is the constant for the mime type "application/lgr+xml".
	ApplicationLgrXml = "application/lgr+xml"

	// ApplicationLinkFormat is the constant for the mime type "application/link-format".
	ApplicationLinkFormat = "application/link-format"

	// ApplicationLoadControlXml is the constant for the mime type "application/load-control+xml".
	ApplicationLoadControlXml = "application/load-control+xml"

	// ApplicationLostXml is the constant for the mime type "application/lost+xml".
	ApplicationLostXml = "application/lost+xml"

	// ApplicationLostsyncXml is the constant for the mime type "application/lostsync+xml".
	ApplicationLostsyncXml = "application/lostsync+xml"

	// ApplicationLXF is the constant for the mime type "application/LXF".
	ApplicationLXF = "application/LXF"

	// ApplicationMacBinhex40 is the constant for the mime type "application/mac-binhex40".
	ApplicationMacBinhex40 = "application/mac-binhex40"

	// ApplicationMacwriteii is the constant for the mime type "application/macwriteii".
	ApplicationMacwriteii = "application/macwriteii"

	// ApplicationMadsXml is the constant for the mime type "application/mads+xml".
	ApplicationMadsXml = "application/mads+xml"

	// ApplicationMarc is the constant for the mime type "application/marc".
	ApplicationMarc = "application/marc"

	// ApplicationMarcxmlXml is the constant for the mime type "application/marcxml+xml".
	ApplicationMarcxmlXml = "application/marcxml+xml"

	// ApplicationMathematica is the constant for the mime type "application/mathematica".
	ApplicationMathematica = "application/mathematica"

	// ApplicationMathmlContentXml is the constant for the mime type "application/mathml-content+xml".
	ApplicationMathmlContentXml = "application/mathml-content+xml"

	// ApplicationMathmlPresentationXml is the constant for the mime type "application/mathml-presentation+xml".
	ApplicationMathmlPresentationXml = "application/mathml-presentation+xml"

	// ApplicationMathmlXml is the constant for the mime type "application/mathml+xml".
	ApplicationMathmlXml = "application/mathml+xml"

	// ApplicationMbmsAssociatedProcedureDescriptionXml is the constant for the mime type "application/mbms-associated-procedure-description+xml".
	ApplicationMbmsAssociatedProcedureDescriptionXml = "application/mbms-associated-procedure-description+xml"

	// ApplicationMbmsDeregisterXml is the constant for the mime type "application/mbms-deregister+xml".
	ApplicationMbmsDeregisterXml = "application/mbms-deregister+xml"

	// ApplicationMbmsEnvelopeXml is the constant for the mime type "application/mbms-envelope+xml".
	ApplicationMbmsEnvelopeXml = "application/mbms-envelope+xml"

	// ApplicationMbmsMskResponseXml is the constant for the mime type "application/mbms-msk-response+xml".
	ApplicationMbmsMskResponseXml = "application/mbms-msk-response+xml"

	// ApplicationMbmsMskXml is the constant for the mime type "application/mbms-msk+xml".
	ApplicationMbmsMskXml = "application/mbms-msk+xml"

	// ApplicationMbmsProtectionDescriptionXml is the constant for the mime type "application/mbms-protection-description+xml".
	ApplicationMbmsProtectionDescriptionXml = "application/mbms-protection-description+xml"

	// ApplicationMbmsReceptionReportXml is the constant for the mime type "application/mbms-reception-report+xml".
	ApplicationMbmsReceptionReportXml = "application/mbms-reception-report+xml"

	// ApplicationMbmsRegisterResponseXml is the constant for the mime type "application/mbms-register-response+xml".
	ApplicationMbmsRegisterResponseXml = "application/mbms-register-response+xml"

	// ApplicationMbmsRegisterXml is the constant for the mime type "application/mbms-register+xml".
	ApplicationMbmsRegisterXml = "application/mbms-register+xml"

	// ApplicationMbmsScheduleXml is the constant for the mime type "application/mbms-schedule+xml".
	ApplicationMbmsScheduleXml = "application/mbms-schedule+xml"

	// ApplicationMbmsUserServiceDescriptionXml is the constant for the mime type "application/mbms-user-service-description+xml".
	ApplicationMbmsUserServiceDescriptionXml = "application/mbms-user-service-description+xml"

	// ApplicationMbox is the constant for the mime type "application/mbox".
	ApplicationMbox = "application/mbox"

	// ApplicationMediaControlXml is the constant for the mime type "application/media_control+xml".
	ApplicationMediaControlXml = "application/media_control+xml"

	// ApplicationMediaPolicyDatasetXml is the constant for the mime type "application/media-policy-dataset+xml".
	ApplicationMediaPolicyDatasetXml = "application/media-policy-dataset+xml"

	// ApplicationMediaservercontrolXml is the constant for the mime type "application/mediaservercontrol+xml".
	ApplicationMediaservercontrolXml = "application/mediaservercontrol+xml"

	// ApplicationMergePatchJson is the constant for the mime type "application/merge-patch+json".
	ApplicationMergePatchJson = "application/merge-patch+json"

	// ApplicationMetalink4Xml is the constant for the mime type "application/metalink4+xml".
	ApplicationMetalink4Xml = "application/metalink4+xml"

	// ApplicationMetsXml is the constant for the mime type "application/mets+xml".
	ApplicationMetsXml = "application/mets+xml"

	// ApplicationMF4 is the constant for the mime type "application/MF4".
	ApplicationMF4 = "application/MF4"

	// ApplicationMikey is the constant for the mime type "application/mikey".
	ApplicationMikey = "application/mikey"

	// ApplicationMipc is the constant for the mime type "application/mipc".
	ApplicationMipc = "application/mipc"

	// ApplicationMmtAeiXml is the constant for the mime type "application/mmt-aei+xml".
	ApplicationMmtAeiXml = "application/mmt-aei+xml"

	// ApplicationMmtUsdXml is the constant for the mime type "application/mmt-usd+xml".
	ApplicationMmtUsdXml = "application/mmt-usd+xml"

	// ApplicationModsXml is the constant for the mime type "application/mods+xml".
	ApplicationModsXml = "application/mods+xml"

	// ApplicationMossKeys is the constant for the mime type "application/moss-keys".
	ApplicationMossKeys = "application/moss-keys"

	// ApplicationMossSignature is the constant for the mime type "application/moss-signature".
	ApplicationMossSignature = "application/moss-signature"

	// ApplicationMosskeyData is the constant for the mime type "application/mosskey-data".
	ApplicationMosskeyData = "application/mosskey-data"

	// ApplicationMosskeyRequest is the constant for the mime type "application/mosskey-request".
	ApplicationMosskeyRequest = "application/mosskey-request"

	// ApplicationMp21 is the constant for the mime type "application/mp21".
	ApplicationMp21 = "application/mp21"

	// ApplicationMp4 is the constant for the mime type "application/mp4".
	ApplicationMp4 = "application/mp4"

	// ApplicationMpeg4Generic is the constant for the mime type "application/mpeg4-generic".
	ApplicationMpeg4Generic = "application/mpeg4-generic"

	// ApplicationMpeg4Iod is the constant for the mime type "application/mpeg4-iod".
	ApplicationMpeg4Iod = "application/mpeg4-iod"

	// ApplicationMpeg4IodXmt is the constant for the mime type "application/mpeg4-iod-xmt".
	ApplicationMpeg4IodXmt = "application/mpeg4-iod-xmt"

	// ApplicationMrbConsumerXml is the constant for the mime type "application/mrb-consumer+xml".
	ApplicationMrbConsumerXml = "application/mrb-consumer+xml"

	// ApplicationMrbPublishXml is the constant for the mime type "application/mrb-publish+xml".
	ApplicationMrbPublishXml = "application/mrb-publish+xml"

	// ApplicationMscIvrXml is the constant for the mime type "application/msc-ivr+xml".
	ApplicationMscIvrXml = "application/msc-ivr+xml"

	// ApplicationMscMixerXml is the constant for the mime type "application/msc-mixer+xml".
	ApplicationMscMixerXml = "application/msc-mixer+xml"

	// ApplicationMsword is the constant for the mime type "application/msword".
	ApplicationMsword = "application/msword"

	// ApplicationMudJson is the constant for the mime type "application/mud+json".
	ApplicationMudJson = "application/mud+json"

	// ApplicationMxf is the constant for the mime type "application/mxf".
	ApplicationMxf = "application/mxf"

	// ApplicationNQuads is the constant for the mime type "application/n-quads".
	ApplicationNQuads = "application/n-quads"

	// ApplicationNTriples is the constant for the mime type "application/n-triples".
	ApplicationNTriples = "application/n-triples"

	// ApplicationNasdata is the constant for the mime type "application/nasdata".
	ApplicationNasdata = "application/nasdata"

	// ApplicationNewsCheckgroups is the constant for the mime type "application/news-checkgroups".
	ApplicationNewsCheckgroups = "application/news-checkgroups"

	// ApplicationNewsGroupinfo is the constant for the mime type "application/news-groupinfo".
	ApplicationNewsGroupinfo = "application/news-groupinfo"

	// ApplicationNewsTransmission is the constant for the mime type "application/news-transmission".
	ApplicationNewsTransmission = "application/news-transmission"

	// ApplicationNlsmlXml is the constant for the mime type "application/nlsml+xml".
	ApplicationNlsmlXml = "application/nlsml+xml"

	// ApplicationNode is the constant for the mime type "application/node".
	ApplicationNode = "application/node"

	// ApplicationNss is the constant for the mime type "application/nss".
	ApplicationNss = "application/nss"

	// ApplicationOcspRequest is the constant for the mime type "application/ocsp-request".
	ApplicationOcspRequest = "application/ocsp-request"

	// ApplicationOcspResponse is the constant for the mime type "application/ocsp-response".
	ApplicationOcspResponse = "application/ocsp-response"

	// ApplicationOctetStream is the constant for the mime type "application/octet-stream".
	ApplicationOctetStream = "application/octet-stream"

	// ApplicationODA is the constant for the mime type "application/ODA".
	ApplicationODA = "application/ODA"

	// ApplicationOdmXml is the constant for the mime type "application/odm+xml".
	ApplicationOdmXml = "application/odm+xml"

	// ApplicationODX is the constant for the mime type "application/ODX".
	ApplicationODX = "application/ODX"

	// ApplicationOebpsPackageXml is the constant for the mime type "application/oebps-package+xml".
	ApplicationOebpsPackageXml = "application/oebps-package+xml"

	// ApplicationOgg is the constant for the mime type "application/ogg".
	ApplicationOgg = "application/ogg"

	// ApplicationOscore is the constant for the mime type "application/oscore".
	ApplicationOscore = "application/oscore"

	// ApplicationOxps is the constant for the mime type "application/oxps".
	ApplicationOxps = "application/oxps"

	// ApplicationP2POverlayXml is the constant for the mime type "application/p2p-overlay+xml".
	ApplicationP2POverlayXml = "application/p2p-overlay+xml"

	// ApplicationParityfec is the constant for the mime type "application/parityfec".
	ApplicationParityfec = "application/parityfec"

	// ApplicationPassport is the constant for the mime type "application/passport".
	ApplicationPassport = "application/passport"

	// ApplicationPatchOpsErrorXml is the constant for the mime type "application/patch-ops-error+xml".
	ApplicationPatchOpsErrorXml = "application/patch-ops-error+xml"

	// ApplicationPdf is the constant for the mime type "application/pdf".
	ApplicationPdf = "application/pdf"

	// ApplicationPDX is the constant for the mime type "application/PDX".
	ApplicationPDX = "application/PDX"

	// ApplicationPemCertificateChain is the constant for the mime type "application/pem-certificate-chain".
	ApplicationPemCertificateChain = "application/pem-certificate-chain"

	// ApplicationPgpEncrypted is the constant for the mime type "application/pgp-encrypted".
	ApplicationPgpEncrypted = "application/pgp-encrypted"

	// ApplicationPgpKeys is the constant for the mime type "application/pgp-keys".
	ApplicationPgpKeys = "application/pgp-keys"

	// ApplicationPgpSignature is the constant for the mime type "application/pgp-signature".
	ApplicationPgpSignature = "application/pgp-signature"

	// ApplicationPidfDiffXml is the constant for the mime type "application/pidf-diff+xml".
	ApplicationPidfDiffXml = "application/pidf-diff+xml"

	// ApplicationPidfXml is the constant for the mime type "application/pidf+xml".
	ApplicationPidfXml = "application/pidf+xml"

	// ApplicationPkcs10 is the constant for the mime type "application/pkcs10".
	ApplicationPkcs10 = "application/pkcs10"

	// ApplicationPkcs7Mime is the constant for the mime type "application/pkcs7-mime".
	ApplicationPkcs7Mime = "application/pkcs7-mime"

	// ApplicationPkcs7Signature is the constant for the mime type "application/pkcs7-signature".
	ApplicationPkcs7Signature = "application/pkcs7-signature"

	// ApplicationPkcs8 is the constant for the mime type "application/pkcs8".
	ApplicationPkcs8 = "application/pkcs8"

	// ApplicationPkcs8Encrypted is the constant for the mime type "application/pkcs8-encrypted".
	ApplicationPkcs8Encrypted = "application/pkcs8-encrypted"

	// ApplicationPkcs12 is the constant for the mime type "application/pkcs12".
	ApplicationPkcs12 = "application/pkcs12"

	// ApplicationPkixAttrCert is the constant for the mime type "application/pkix-attr-cert".
	ApplicationPkixAttrCert = "application/pkix-attr-cert"

	// ApplicationPkixCert is the constant for the mime type "application/pkix-cert".
	ApplicationPkixCert = "application/pkix-cert"

	// ApplicationPkixCrl is the constant for the mime type "application/pkix-crl".
	ApplicationPkixCrl = "application/pkix-crl"

	// ApplicationPkixPkipath is the constant for the mime type "application/pkix-pkipath".
	ApplicationPkixPkipath = "application/pkix-pkipath"

	// ApplicationPkixcmp is the constant for the mime type "application/pkixcmp".
	ApplicationPkixcmp = "application/pkixcmp"

	// ApplicationPlsXml is the constant for the mime type "application/pls+xml".
	ApplicationPlsXml = "application/pls+xml"

	// ApplicationPocSettingsXml is the constant for the mime type "application/poc-settings+xml".
	ApplicationPocSettingsXml = "application/poc-settings+xml"

	// ApplicationPostscript is the constant for the mime type "application/postscript".
	ApplicationPostscript = "application/postscript"

	// ApplicationPpspTrackerJson is the constant for the mime type "application/ppsp-tracker+json".
	ApplicationPpspTrackerJson = "application/ppsp-tracker+json"

	// ApplicationProblemJson is the constant for the mime type "application/problem+json".
	ApplicationProblemJson = "application/problem+json"

	// ApplicationProblemXml is the constant for the mime type "application/problem+xml".
	ApplicationProblemXml = "application/problem+xml"

	// ApplicationProvenanceXml is the constant for the mime type "application/provenance+xml".
	ApplicationProvenanceXml = "application/provenance+xml"

	// ApplicationPrsAlvestrandTitraxSheet is the constant for the mime type "application/prs.alvestrand.titrax-sheet".
	ApplicationPrsAlvestrandTitraxSheet = "application/prs.alvestrand.titrax-sheet"

	// ApplicationPrsCww is the constant for the mime type "application/prs.cww".
	ApplicationPrsCww = "application/prs.cww"

	// ApplicationPrsHpubZip is the constant for the mime type "application/prs.hpub+zip".
	ApplicationPrsHpubZip = "application/prs.hpub+zip"

	// ApplicationPrsNprend is the constant for the mime type "application/prs.nprend".
	ApplicationPrsNprend = "application/prs.nprend"

	// ApplicationPrsPlucker is the constant for the mime type "application/prs.plucker".
	ApplicationPrsPlucker = "application/prs.plucker"

	// ApplicationPrsRdfXmlCrypt is the constant for the mime type "application/prs.rdf-xml-crypt".
	ApplicationPrsRdfXmlCrypt = "application/prs.rdf-xml-crypt"

	// ApplicationPrsXsfXml is the constant for the mime type "application/prs.xsf+xml".
	ApplicationPrsXsfXml = "application/prs.xsf+xml"

	// ApplicationPskcXml is the constant for the mime type "application/pskc+xml".
	ApplicationPskcXml = "application/pskc+xml"

	// ApplicationRdfXml is the constant for the mime type "application/rdf+xml".
	ApplicationRdfXml = "application/rdf+xml"

	// ApplicationRouteApdXml is the constant for the mime type "application/route-apd+xml".
	ApplicationRouteApdXml = "application/route-apd+xml"

	// ApplicationRouteSTsidXml is the constant for the mime type "application/route-s-tsid+xml".
	ApplicationRouteSTsidXml = "application/route-s-tsid+xml"

	// ApplicationRouteUsdXml is the constant for the mime type "application/route-usd+xml".
	ApplicationRouteUsdXml = "application/route-usd+xml"

	// ApplicationQSIG is the constant for the mime type "application/QSIG".
	ApplicationQSIG = "application/QSIG"

	// ApplicationRaptorfec is the constant for the mime type "application/raptorfec".
	ApplicationRaptorfec = "application/raptorfec"

	// ApplicationRdapJson is the constant for the mime type "application/rdap+json".
	ApplicationRdapJson = "application/rdap+json"

	// ApplicationReginfoXml is the constant for the mime type "application/reginfo+xml".
	ApplicationReginfoXml = "application/reginfo+xml"

	// ApplicationRelaxNgCompactSyntax is the constant for the mime type "application/relax-ng-compact-syntax".
	ApplicationRelaxNgCompactSyntax = "application/relax-ng-compact-syntax"

	// ApplicationRemotePrinting is the constant for the mime type "application/remote-printing".
	ApplicationRemotePrinting = "application/remote-printing"

	// ApplicationReputonJson is the constant for the mime type "application/reputon+json".
	ApplicationReputonJson = "application/reputon+json"

	// ApplicationResourceListsDiffXml is the constant for the mime type "application/resource-lists-diff+xml".
	ApplicationResourceListsDiffXml = "application/resource-lists-diff+xml"

	// ApplicationResourceListsXml is the constant for the mime type "application/resource-lists+xml".
	ApplicationResourceListsXml = "application/resource-lists+xml"

	// ApplicationRfcXml is the constant for the mime type "application/rfc+xml".
	ApplicationRfcXml = "application/rfc+xml"

	// ApplicationRiscos is the constant for the mime type "application/riscos".
	ApplicationRiscos = "application/riscos"

	// ApplicationRlmiXml is the constant for the mime type "application/rlmi+xml".
	ApplicationRlmiXml = "application/rlmi+xml"

	// ApplicationRlsServicesXml is the constant for the mime type "application/rls-services+xml".
	ApplicationRlsServicesXml = "application/rls-services+xml"

	// ApplicationRpkiGhostbusters is the constant for the mime type "application/rpki-ghostbusters".
	ApplicationRpkiGhostbusters = "application/rpki-ghostbusters"

	// ApplicationRpkiManifest is the constant for the mime type "application/rpki-manifest".
	ApplicationRpkiManifest = "application/rpki-manifest"

	// ApplicationRpkiPublication is the constant for the mime type "application/rpki-publication".
	ApplicationRpkiPublication = "application/rpki-publication"

	// ApplicationRpkiRoa is the constant for the mime type "application/rpki-roa".
	ApplicationRpkiRoa = "application/rpki-roa"

	// ApplicationRpkiUpdown is the constant for the mime type "application/rpki-updown".
	ApplicationRpkiUpdown = "application/rpki-updown"

	// ApplicationRtf is the constant for the mime type "application/rtf".
	ApplicationRtf = "application/rtf"

	// ApplicationRtploopback is the constant for the mime type "application/rtploopback".
	ApplicationRtploopback = "application/rtploopback"

	// ApplicationRtx is the constant for the mime type "application/rtx".
	ApplicationRtx = "application/rtx"

	// ApplicationSamlassertionXml is the constant for the mime type "application/samlassertion+xml".
	ApplicationSamlassertionXml = "application/samlassertion+xml"

	// ApplicationSamlmetadataXml is the constant for the mime type "application/samlmetadata+xml".
	ApplicationSamlmetadataXml = "application/samlmetadata+xml"

	// ApplicationSbmlXml is the constant for the mime type "application/sbml+xml".
	ApplicationSbmlXml = "application/sbml+xml"

	// ApplicationScaipXml is the constant for the mime type "application/scaip+xml".
	ApplicationScaipXml = "application/scaip+xml"

	// ApplicationScimJson is the constant for the mime type "application/scim+json".
	ApplicationScimJson = "application/scim+json"

	// ApplicationScvpCvRequest is the constant for the mime type "application/scvp-cv-request".
	ApplicationScvpCvRequest = "application/scvp-cv-request"

	// ApplicationScvpCvResponse is the constant for the mime type "application/scvp-cv-response".
	ApplicationScvpCvResponse = "application/scvp-cv-response"

	// ApplicationScvpVpRequest is the constant for the mime type "application/scvp-vp-request".
	ApplicationScvpVpRequest = "application/scvp-vp-request"

	// ApplicationScvpVpResponse is the constant for the mime type "application/scvp-vp-response".
	ApplicationScvpVpResponse = "application/scvp-vp-response"

	// ApplicationSdp is the constant for the mime type "application/sdp".
	ApplicationSdp = "application/sdp"

	// ApplicationSeceventJwt is the constant for the mime type "application/secevent+jwt".
	ApplicationSeceventJwt = "application/secevent+jwt"

	// ApplicationSenmlExi is the constant for the mime type "application/senml-exi".
	ApplicationSenmlExi = "application/senml-exi"

	// ApplicationSenmlCbor is the constant for the mime type "application/senml+cbor".
	ApplicationSenmlCbor = "application/senml+cbor"

	// ApplicationSenmlJson is the constant for the mime type "application/senml+json".
	ApplicationSenmlJson = "application/senml+json"

	// ApplicationSenmlXml is the constant for the mime type "application/senml+xml".
	ApplicationSenmlXml = "application/senml+xml"

	// ApplicationSensmlExi is the constant for the mime type "application/sensml-exi".
	ApplicationSensmlExi = "application/sensml-exi"

	// ApplicationSensmlCbor is the constant for the mime type "application/sensml+cbor".
	ApplicationSensmlCbor = "application/sensml+cbor"

	// ApplicationSensmlJson is the constant for the mime type "application/sensml+json".
	ApplicationSensmlJson = "application/sensml+json"

	// ApplicationSensmlXml is the constant for the mime type "application/sensml+xml".
	ApplicationSensmlXml = "application/sensml+xml"

	// ApplicationSepExi is the constant for the mime type "application/sep-exi".
	ApplicationSepExi = "application/sep-exi"

	// ApplicationSepXml is the constant for the mime type "application/sep+xml".
	ApplicationSepXml = "application/sep+xml"

	// ApplicationSessionInfo is the constant for the mime type "application/session-info".
	ApplicationSessionInfo = "application/session-info"

	// ApplicationSetPayment is the constant for the mime type "application/set-payment".
	ApplicationSetPayment = "application/set-payment"

	// ApplicationSetPaymentInitiation is the constant for the mime type "application/set-payment-initiation".
	ApplicationSetPaymentInitiation = "application/set-payment-initiation"

	// ApplicationSetRegistration is the constant for the mime type "application/set-registration".
	ApplicationSetRegistration = "application/set-registration"

	// ApplicationSetRegistrationInitiation is the constant for the mime type "application/set-registration-initiation".
	ApplicationSetRegistrationInitiation = "application/set-registration-initiation"

	// ApplicationSgml is the constant for the mime type "application/sgml".
	ApplicationSgml = "application/sgml"

	// ApplicationSgmlOpenCatalog is the constant for the mime type "application/sgml-open-catalog".
	ApplicationSgmlOpenCatalog = "application/sgml-open-catalog"

	// ApplicationShfXml is the constant for the mime type "application/shf+xml".
	ApplicationShfXml = "application/shf+xml"

	// ApplicationSieve is the constant for the mime type "application/sieve".
	ApplicationSieve = "application/sieve"

	// ApplicationSimpleFilterXml is the constant for the mime type "application/simple-filter+xml".
	ApplicationSimpleFilterXml = "application/simple-filter+xml"

	// ApplicationSimpleMessageSummary is the constant for the mime type "application/simple-message-summary".
	ApplicationSimpleMessageSummary = "application/simple-message-summary"

	// ApplicationSimpleSymbolContainer is the constant for the mime type "application/simpleSymbolContainer".
	ApplicationSimpleSymbolContainer = "application/simpleSymbolContainer"

	// ApplicationSipc is the constant for the mime type "application/sipc".
	ApplicationSipc = "application/sipc"

	// ApplicationSlate is the constant for the mime type "application/slate".
	ApplicationSlate = "application/slate"

	// ApplicationSmil is the constant for the mime type "application/smil - OBSOLETED in favor of application/smil+xml".
	ApplicationSmil = "application/smil"

	// ApplicationSmilXml is the constant for the mime type "application/smil+xml".
	ApplicationSmilXml = "application/smil+xml"

	// ApplicationSmpte336M is the constant for the mime type "application/smpte336m".
	ApplicationSmpte336M = "application/smpte336m"

	// ApplicationSoapFastinfoset is the constant for the mime type "application/soap+fastinfoset".
	ApplicationSoapFastinfoset = "application/soap+fastinfoset"

	// ApplicationSoapXml is the constant for the mime type "application/soap+xml".
	ApplicationSoapXml = "application/soap+xml"

	// ApplicationSparqlQuery is the constant for the mime type "application/sparql-query".
	ApplicationSparqlQuery = "application/sparql-query"

	// ApplicationSparqlResultsXml is the constant for the mime type "application/sparql-results+xml".
	ApplicationSparqlResultsXml = "application/sparql-results+xml"

	// ApplicationSpiritsEventXml is the constant for the mime type "application/spirits-event+xml".
	ApplicationSpiritsEventXml = "application/spirits-event+xml"

	// ApplicationSql is the constant for the mime type "application/sql".
	ApplicationSql = "application/sql"

	// ApplicationSrgs is the constant for the mime type "application/srgs".
	ApplicationSrgs = "application/srgs"

	// ApplicationSrgsXml is the constant for the mime type "application/srgs+xml".
	ApplicationSrgsXml = "application/srgs+xml"

	// ApplicationSruXml is the constant for the mime type "application/sru+xml".
	ApplicationSruXml = "application/sru+xml"

	// ApplicationSsmlXml is the constant for the mime type "application/ssml+xml".
	ApplicationSsmlXml = "application/ssml+xml"

	// ApplicationStixJson is the constant for the mime type "application/stix+json".
	ApplicationStixJson = "application/stix+json"

	// ApplicationSwidXml is the constant for the mime type "application/swid+xml".
	ApplicationSwidXml = "application/swid+xml"

	// ApplicationTampApexUpdate is the constant for the mime type "application/tamp-apex-update".
	ApplicationTampApexUpdate = "application/tamp-apex-update"

	// ApplicationTampApexUpdateConfirm is the constant for the mime type "application/tamp-apex-update-confirm".
	ApplicationTampApexUpdateConfirm = "application/tamp-apex-update-confirm"

	// ApplicationTampCommunityUpdate is the constant for the mime type "application/tamp-community-update".
	ApplicationTampCommunityUpdate = "application/tamp-community-update"

	// ApplicationTampCommunityUpdateConfirm is the constant for the mime type "application/tamp-community-update-confirm".
	ApplicationTampCommunityUpdateConfirm = "application/tamp-community-update-confirm"

	// ApplicationTampError is the constant for the mime type "application/tamp-error".
	ApplicationTampError = "application/tamp-error"

	// ApplicationTampSequenceAdjust is the constant for the mime type "application/tamp-sequence-adjust".
	ApplicationTampSequenceAdjust = "application/tamp-sequence-adjust"

	// ApplicationTampSequenceAdjustConfirm is the constant for the mime type "application/tamp-sequence-adjust-confirm".
	ApplicationTampSequenceAdjustConfirm = "application/tamp-sequence-adjust-confirm"

	// ApplicationTampStatusQuery is the constant for the mime type "application/tamp-status-query".
	ApplicationTampStatusQuery = "application/tamp-status-query"

	// ApplicationTampStatusResponse is the constant for the mime type "application/tamp-status-response".
	ApplicationTampStatusResponse = "application/tamp-status-response"

	// ApplicationTampUpdate is the constant for the mime type "application/tamp-update".
	ApplicationTampUpdate = "application/tamp-update"

	// ApplicationTampUpdateConfirm is the constant for the mime type "application/tamp-update-confirm".
	ApplicationTampUpdateConfirm = "application/tamp-update-confirm"

	// ApplicationTaxiiJson is the constant for the mime type "application/taxii+json".
	ApplicationTaxiiJson = "application/taxii+json"

	// ApplicationTeiXml is the constant for the mime type "application/tei+xml".
	ApplicationTeiXml = "application/tei+xml"

	// ApplicationTETRAISI is the constant for the mime type "application/TETRA_ISI".
	ApplicationTETRAISI = "application/TETRA_ISI"

	// ApplicationThraudXml is the constant for the mime type "application/thraud+xml".
	ApplicationThraudXml = "application/thraud+xml"

	// ApplicationTimestampQuery is the constant for the mime type "application/timestamp-query".
	ApplicationTimestampQuery = "application/timestamp-query"

	// ApplicationTimestampReply is the constant for the mime type "application/timestamp-reply".
	ApplicationTimestampReply = "application/timestamp-reply"

	// ApplicationTimestampedData is the constant for the mime type "application/timestamped-data".
	ApplicationTimestampedData = "application/timestamped-data"

	// ApplicationTlsrptGzip is the constant for the mime type "application/tlsrpt+gzip".
	ApplicationTlsrptGzip = "application/tlsrpt+gzip"

	// ApplicationTlsrptJson is the constant for the mime type "application/tlsrpt+json".
	ApplicationTlsrptJson = "application/tlsrpt+json"

	// ApplicationTnauthlist is the constant for the mime type "application/tnauthlist".
	ApplicationTnauthlist = "application/tnauthlist"

	// ApplicationTrickleIceSdpfrag is the constant for the mime type "application/trickle-ice-sdpfrag".
	ApplicationTrickleIceSdpfrag = "application/trickle-ice-sdpfrag"

	// ApplicationTrig is the constant for the mime type "application/trig".
	ApplicationTrig = "application/trig"

	// ApplicationTtmlXml is the constant for the mime type "application/ttml+xml".
	ApplicationTtmlXml = "application/ttml+xml"

	// ApplicationTveTrigger is the constant for the mime type "application/tve-trigger".
	ApplicationTveTrigger = "application/tve-trigger"

	// ApplicationTzif is the constant for the mime type "application/tzif".
	ApplicationTzif = "application/tzif"

	// ApplicationTzifLeap is the constant for the mime type "application/tzif-leap".
	ApplicationTzifLeap = "application/tzif-leap"

	// ApplicationUlpfec is the constant for the mime type "application/ulpfec".
	ApplicationUlpfec = "application/ulpfec"

	// ApplicationUrcGrpsheetXml is the constant for the mime type "application/urc-grpsheet+xml".
	ApplicationUrcGrpsheetXml = "application/urc-grpsheet+xml"

	// ApplicationUrcRessheetXml is the constant for the mime type "application/urc-ressheet+xml".
	ApplicationUrcRessheetXml = "application/urc-ressheet+xml"

	// ApplicationUrcTargetdescXml is the constant for the mime type "application/urc-targetdesc+xml".
	ApplicationUrcTargetdescXml = "application/urc-targetdesc+xml"

	// ApplicationUrcUisocketdescXml is the constant for the mime type "application/urc-uisocketdesc+xml".
	ApplicationUrcUisocketdescXml = "application/urc-uisocketdesc+xml"

	// ApplicationVcardJson is the constant for the mime type "application/vcard+json".
	ApplicationVcardJson = "application/vcard+json"

	// ApplicationVcardXml is the constant for the mime type "application/vcard+xml".
	ApplicationVcardXml = "application/vcard+xml"

	// ApplicationVemmi is the constant for the mime type "application/vemmi".
	ApplicationVemmi = "application/vemmi"

	// ApplicationVnd1000mindsDecisionModelXml is the constant for the mime type "application/vnd.1000minds.decision-model+xml".
	ApplicationVnd1000mindsDecisionModelXml = "application/vnd.1000minds.decision-model+xml"

	// ApplicationVnd3gppAccessTransferEventsXml is the constant for the mime type "application/vnd.3gpp.access-transfer-events+xml".
	ApplicationVnd3gppAccessTransferEventsXml = "application/vnd.3gpp.access-transfer-events+xml"

	// ApplicationVnd3gppBsfXml is the constant for the mime type "application/vnd.3gpp.bsf+xml".
	ApplicationVnd3gppBsfXml = "application/vnd.3gpp.bsf+xml"

	// ApplicationVnd3gppGMOPXml is the constant for the mime type "application/vnd.3gpp.GMOP+xml".
	ApplicationVnd3gppGMOPXml = "application/vnd.3gpp.GMOP+xml"

	// ApplicationVnd3gppMcSignallingEar is the constant for the mime type "application/vnd.3gpp.mc-signalling-ear".
	ApplicationVnd3gppMcSignallingEar = "application/vnd.3gpp.mc-signalling-ear"

	// ApplicationVnd3gppMcdataAffiliationCommandXml is the constant for the mime type "application/vnd.3gpp.mcdata-affiliation-command+xml".
	ApplicationVnd3gppMcdataAffiliationCommandXml = "application/vnd.3gpp.mcdata-affiliation-command+xml"

	// ApplicationVnd3gppMcdataInfoXml is the constant for the mime type "application/vnd.3gpp.mcdata-info+xml".
	ApplicationVnd3gppMcdataInfoXml = "application/vnd.3gpp.mcdata-info+xml"

	// ApplicationVnd3gppMcdataPayload is the constant for the mime type "application/vnd.3gpp.mcdata-payload".
	ApplicationVnd3gppMcdataPayload = "application/vnd.3gpp.mcdata-payload"

	// ApplicationVnd3gppMcdataServiceConfigXml is the constant for the mime type "application/vnd.3gpp.mcdata-service-config+xml".
	ApplicationVnd3gppMcdataServiceConfigXml = "application/vnd.3gpp.mcdata-service-config+xml"

	// ApplicationVnd3gppMcdataSignalling is the constant for the mime type "application/vnd.3gpp.mcdata-signalling".
	ApplicationVnd3gppMcdataSignalling = "application/vnd.3gpp.mcdata-signalling"

	// ApplicationVnd3gppMcdataUeConfigXml is the constant for the mime type "application/vnd.3gpp.mcdata-ue-config+xml".
	ApplicationVnd3gppMcdataUeConfigXml = "application/vnd.3gpp.mcdata-ue-config+xml"

	// ApplicationVnd3gppMcdataUserProfileXml is the constant for the mime type "application/vnd.3gpp.mcdata-user-profile+xml".
	ApplicationVnd3gppMcdataUserProfileXml = "application/vnd.3gpp.mcdata-user-profile+xml"

	// ApplicationVnd3gppMcpttAffiliationCommandXml is the constant for the mime type "application/vnd.3gpp.mcptt-affiliation-command+xml".
	ApplicationVnd3gppMcpttAffiliationCommandXml = "application/vnd.3gpp.mcptt-affiliation-command+xml"

	// ApplicationVnd3gppMcpttFloorRequestXml is the constant for the mime type "application/vnd.3gpp.mcptt-floor-request+xml".
	ApplicationVnd3gppMcpttFloorRequestXml = "application/vnd.3gpp.mcptt-floor-request+xml"

	// ApplicationVnd3gppMcpttInfoXml is the constant for the mime type "application/vnd.3gpp.mcptt-info+xml".
	ApplicationVnd3gppMcpttInfoXml = "application/vnd.3gpp.mcptt-info+xml"

	// ApplicationVnd3gppMcpttLocationInfoXml is the constant for the mime type "application/vnd.3gpp.mcptt-location-info+xml".
	ApplicationVnd3gppMcpttLocationInfoXml = "application/vnd.3gpp.mcptt-location-info+xml"

	// ApplicationVnd3gppMcpttMbmsUsageInfoXml is the constant for the mime type "application/vnd.3gpp.mcptt-mbms-usage-info+xml".
	ApplicationVnd3gppMcpttMbmsUsageInfoXml = "application/vnd.3gpp.mcptt-mbms-usage-info+xml"

	// ApplicationVnd3gppMcpttServiceConfigXml is the constant for the mime type "application/vnd.3gpp.mcptt-service-config+xml".
	ApplicationVnd3gppMcpttServiceConfigXml = "application/vnd.3gpp.mcptt-service-config+xml"

	// ApplicationVnd3gppMcpttSignedXml is the constant for the mime type "application/vnd.3gpp.mcptt-signed+xml".
	ApplicationVnd3gppMcpttSignedXml = "application/vnd.3gpp.mcptt-signed+xml"

	// ApplicationVnd3gppMcpttUeConfigXml is the constant for the mime type "application/vnd.3gpp.mcptt-ue-config+xml".
	ApplicationVnd3gppMcpttUeConfigXml = "application/vnd.3gpp.mcptt-ue-config+xml"

	// ApplicationVnd3gppMcpttUeInitConfigXml is the constant for the mime type "application/vnd.3gpp.mcptt-ue-init-config+xml".
	ApplicationVnd3gppMcpttUeInitConfigXml = "application/vnd.3gpp.mcptt-ue-init-config+xml"

	// ApplicationVnd3gppMcpttUserProfileXml is the constant for the mime type "application/vnd.3gpp.mcptt-user-profile+xml".
	ApplicationVnd3gppMcpttUserProfileXml = "application/vnd.3gpp.mcptt-user-profile+xml"

	// ApplicationVnd3gppMcvideoAffiliationCommandXml is the constant for the mime type "application/vnd.3gpp.mcvideo-affiliation-command+xml".
	ApplicationVnd3gppMcvideoAffiliationCommandXml = "application/vnd.3gpp.mcvideo-affiliation-command+xml"

	// ApplicationVnd3gppMcvideoAffiliationInfoXml is the constant for the mime type "application/vnd.3gpp.mcvideo-affiliation-info+xml - OBSOLETED in favor of application/vnd.3gpp.mcvideo-info+xml".
	ApplicationVnd3gppMcvideoAffiliationInfoXml = "application/vnd.3gpp.mcvideo-affiliation-info+xml"

	// ApplicationVnd3gppMcvideoInfoXml is the constant for the mime type "application/vnd.3gpp.mcvideo-info+xml".
	ApplicationVnd3gppMcvideoInfoXml = "application/vnd.3gpp.mcvideo-info+xml"

	// ApplicationVnd3gppMcvideoLocationInfoXml is the constant for the mime type "application/vnd.3gpp.mcvideo-location-info+xml".
	ApplicationVnd3gppMcvideoLocationInfoXml = "application/vnd.3gpp.mcvideo-location-info+xml"

	// ApplicationVnd3gppMcvideoMbmsUsageInfoXml is the constant for the mime type "application/vnd.3gpp.mcvideo-mbms-usage-info+xml".
	ApplicationVnd3gppMcvideoMbmsUsageInfoXml = "application/vnd.3gpp.mcvideo-mbms-usage-info+xml"

	// ApplicationVnd3gppMcvideoServiceConfigXml is the constant for the mime type "application/vnd.3gpp.mcvideo-service-config+xml".
	ApplicationVnd3gppMcvideoServiceConfigXml = "application/vnd.3gpp.mcvideo-service-config+xml"

	// ApplicationVnd3gppMcvideoTransmissionRequestXml is the constant for the mime type "application/vnd.3gpp.mcvideo-transmission-request+xml".
	ApplicationVnd3gppMcvideoTransmissionRequestXml = "application/vnd.3gpp.mcvideo-transmission-request+xml"

	// ApplicationVnd3gppMcvideoUeConfigXml is the constant for the mime type "application/vnd.3gpp.mcvideo-ue-config+xml".
	ApplicationVnd3gppMcvideoUeConfigXml = "application/vnd.3gpp.mcvideo-ue-config+xml"

	// ApplicationVnd3gppMcvideoUserProfileXml is the constant for the mime type "application/vnd.3gpp.mcvideo-user-profile+xml".
	ApplicationVnd3gppMcvideoUserProfileXml = "application/vnd.3gpp.mcvideo-user-profile+xml"

	// ApplicationVnd3gppMidCallXml is the constant for the mime type "application/vnd.3gpp.mid-call+xml".
	ApplicationVnd3gppMidCallXml = "application/vnd.3gpp.mid-call+xml"

	// ApplicationVnd3gppPicBwLarge is the constant for the mime type "application/vnd.3gpp.pic-bw-large".
	ApplicationVnd3gppPicBwLarge = "application/vnd.3gpp.pic-bw-large"

	// ApplicationVnd3gppPicBwSmall is the constant for the mime type "application/vnd.3gpp.pic-bw-small".
	ApplicationVnd3gppPicBwSmall = "application/vnd.3gpp.pic-bw-small"

	// ApplicationVnd3gppPicBwVar is the constant for the mime type "application/vnd.3gpp.pic-bw-var".
	ApplicationVnd3gppPicBwVar = "application/vnd.3gpp.pic-bw-var"

	// ApplicationVnd3gppProsePc3ChXml is the constant for the mime type "application/vnd.3gpp-prose-pc3ch+xml".
	ApplicationVnd3gppProsePc3ChXml = "application/vnd.3gpp-prose-pc3ch+xml"

	// ApplicationVnd3gppProseXml is the constant for the mime type "application/vnd.3gpp-prose+xml".
	ApplicationVnd3gppProseXml = "application/vnd.3gpp-prose+xml"

	// ApplicationVnd3gppSms is the constant for the mime type "application/vnd.3gpp.sms".
	ApplicationVnd3gppSms = "application/vnd.3gpp.sms"

	// ApplicationVnd3gppSmsXml is the constant for the mime type "application/vnd.3gpp.sms+xml".
	ApplicationVnd3gppSmsXml = "application/vnd.3gpp.sms+xml"

	// ApplicationVnd3gppSrvccExtXml is the constant for the mime type "application/vnd.3gpp.srvcc-ext+xml".
	ApplicationVnd3gppSrvccExtXml = "application/vnd.3gpp.srvcc-ext+xml"

	// ApplicationVnd3gppSRVCCInfoXml is the constant for the mime type "application/vnd.3gpp.SRVCC-info+xml".
	ApplicationVnd3gppSRVCCInfoXml = "application/vnd.3gpp.SRVCC-info+xml"

	// ApplicationVnd3gppStateAndEventInfoXml is the constant for the mime type "application/vnd.3gpp.state-and-event-info+xml".
	ApplicationVnd3gppStateAndEventInfoXml = "application/vnd.3gpp.state-and-event-info+xml"

	// ApplicationVnd3gppUssdXml is the constant for the mime type "application/vnd.3gpp.ussd+xml".
	ApplicationVnd3gppUssdXml = "application/vnd.3gpp.ussd+xml"

	// ApplicationVnd3gppV2XLocalServiceInformation is the constant for the mime type "application/vnd.3gpp-v2x-local-service-information".
	ApplicationVnd3gppV2XLocalServiceInformation = "application/vnd.3gpp-v2x-local-service-information"

	// ApplicationVnd3gpp2BcmcsinfoXml is the constant for the mime type "application/vnd.3gpp2.bcmcsinfo+xml".
	ApplicationVnd3gpp2BcmcsinfoXml = "application/vnd.3gpp2.bcmcsinfo+xml"

	// ApplicationVnd3gpp2Sms is the constant for the mime type "application/vnd.3gpp2.sms".
	ApplicationVnd3gpp2Sms = "application/vnd.3gpp2.sms"

	// ApplicationVnd3gpp2Tcap is the constant for the mime type "application/vnd.3gpp2.tcap".
	ApplicationVnd3gpp2Tcap = "application/vnd.3gpp2.tcap"

	// ApplicationVnd3lightssoftwareImagescal is the constant for the mime type "application/vnd.3lightssoftware.imagescal".
	ApplicationVnd3lightssoftwareImagescal = "application/vnd.3lightssoftware.imagescal"

	// ApplicationVnd3MPostItNotes is the constant for the mime type "application/vnd.3M.Post-it-Notes".
	ApplicationVnd3MPostItNotes = "application/vnd.3M.Post-it-Notes"

	// ApplicationVndAccpacSimplyAso is the constant for the mime type "application/vnd.accpac.simply.aso".
	ApplicationVndAccpacSimplyAso = "application/vnd.accpac.simply.aso"

	// ApplicationVndAccpacSimplyImp is the constant for the mime type "application/vnd.accpac.simply.imp".
	ApplicationVndAccpacSimplyImp = "application/vnd.accpac.simply.imp"

	// ApplicationVndAcucobol is the constant for the mime type "application/vnd.acucobol".
	ApplicationVndAcucobol = "application/vnd.acucobol"

	// ApplicationVndAcucorp is the constant for the mime type "application/vnd.acucorp".
	ApplicationVndAcucorp = "application/vnd.acucorp"

	// ApplicationVndAdobeFlashMovie is the constant for the mime type "application/vnd.adobe.flash.movie".
	ApplicationVndAdobeFlashMovie = "application/vnd.adobe.flash.movie"

	// ApplicationVndAdobeFormscentralFcdt is the constant for the mime type "application/vnd.adobe.formscentral.fcdt".
	ApplicationVndAdobeFormscentralFcdt = "application/vnd.adobe.formscentral.fcdt"

	// ApplicationVndAdobeFxp is the constant for the mime type "application/vnd.adobe.fxp".
	ApplicationVndAdobeFxp = "application/vnd.adobe.fxp"

	// ApplicationVndAdobePartialUpload is the constant for the mime type "application/vnd.adobe.partial-upload".
	ApplicationVndAdobePartialUpload = "application/vnd.adobe.partial-upload"

	// ApplicationVndAdobeXdpXml is the constant for the mime type "application/vnd.adobe.xdp+xml".
	ApplicationVndAdobeXdpXml = "application/vnd.adobe.xdp+xml"

	// ApplicationVndAdobeXfdf is the constant for the mime type "application/vnd.adobe.xfdf".
	ApplicationVndAdobeXfdf = "application/vnd.adobe.xfdf"

	// ApplicationVndAetherImp is the constant for the mime type "application/vnd.aether.imp".
	ApplicationVndAetherImp = "application/vnd.aether.imp"

	// ApplicationVndAfpcAfplinedata is the constant for the mime type "application/vnd.afpc.afplinedata".
	ApplicationVndAfpcAfplinedata = "application/vnd.afpc.afplinedata"

	// ApplicationVndAfpcModca is the constant for the mime type "application/vnd.afpc.modca".
	ApplicationVndAfpcModca = "application/vnd.afpc.modca"

	// ApplicationVndAhBarcode is the constant for the mime type "application/vnd.ah-barcode".
	ApplicationVndAhBarcode = "application/vnd.ah-barcode"

	// ApplicationVndAheadSpace is the constant for the mime type "application/vnd.ahead.space".
	ApplicationVndAheadSpace = "application/vnd.ahead.space"

	// ApplicationVndAirzipFilesecureAzf is the constant for the mime type "application/vnd.airzip.filesecure.azf".
	ApplicationVndAirzipFilesecureAzf = "application/vnd.airzip.filesecure.azf"

	// ApplicationVndAirzipFilesecureAzs is the constant for the mime type "application/vnd.airzip.filesecure.azs".
	ApplicationVndAirzipFilesecureAzs = "application/vnd.airzip.filesecure.azs"

	// ApplicationVndAmadeusJson is the constant for the mime type "application/vnd.amadeus+json".
	ApplicationVndAmadeusJson = "application/vnd.amadeus+json"

	// ApplicationVndAmazonMobi8Ebook is the constant for the mime type "application/vnd.amazon.mobi8-ebook".
	ApplicationVndAmazonMobi8Ebook = "application/vnd.amazon.mobi8-ebook"

	// ApplicationVndAmericandynamicsAcc is the constant for the mime type "application/vnd.americandynamics.acc".
	ApplicationVndAmericandynamicsAcc = "application/vnd.americandynamics.acc"

	// ApplicationVndAmigaAmi is the constant for the mime type "application/vnd.amiga.ami".
	ApplicationVndAmigaAmi = "application/vnd.amiga.ami"

	// ApplicationVndAmundsenMazeXml is the constant for the mime type "application/vnd.amundsen.maze+xml".
	ApplicationVndAmundsenMazeXml = "application/vnd.amundsen.maze+xml"

	// ApplicationVndAndroidOta is the constant for the mime type "application/vnd.android.ota".
	ApplicationVndAndroidOta = "application/vnd.android.ota"

	// ApplicationVndAnki is the constant for the mime type "application/vnd.anki".
	ApplicationVndAnki = "application/vnd.anki"

	// ApplicationVndAnserWebCertificateIssueInitiation is the constant for the mime type "application/vnd.anser-web-certificate-issue-initiation".
	ApplicationVndAnserWebCertificateIssueInitiation = "application/vnd.anser-web-certificate-issue-initiation"

	// ApplicationVndAntixGameComponent is the constant for the mime type "application/vnd.antix.game-component".
	ApplicationVndAntixGameComponent = "application/vnd.antix.game-component"

	// ApplicationVndApacheThriftBinary is the constant for the mime type "application/vnd.apache.thrift.binary".
	ApplicationVndApacheThriftBinary = "application/vnd.apache.thrift.binary"

	// ApplicationVndApacheThriftCompact is the constant for the mime type "application/vnd.apache.thrift.compact".
	ApplicationVndApacheThriftCompact = "application/vnd.apache.thrift.compact"

	// ApplicationVndApacheThriftJson is the constant for the mime type "application/vnd.apache.thrift.json".
	ApplicationVndApacheThriftJson = "application/vnd.apache.thrift.json"

	// ApplicationVndApiJson is the constant for the mime type "application/vnd.api+json".
	ApplicationVndApiJson = "application/vnd.api+json"

	// ApplicationVndApothekendeReservationJson is the constant for the mime type "application/vnd.apothekende.reservation+json".
	ApplicationVndApothekendeReservationJson = "application/vnd.apothekende.reservation+json"

	// ApplicationVndAppleInstallerXml is the constant for the mime type "application/vnd.apple.installer+xml".
	ApplicationVndAppleInstallerXml = "application/vnd.apple.installer+xml"

	// ApplicationVndAppleKeynote is the constant for the mime type "application/vnd.apple.keynote".
	ApplicationVndAppleKeynote = "application/vnd.apple.keynote"

	// ApplicationVndAppleMpegurl is the constant for the mime type "application/vnd.apple.mpegurl".
	ApplicationVndAppleMpegurl = "application/vnd.apple.mpegurl"

	// ApplicationVndAppleNumbers is the constant for the mime type "application/vnd.apple.numbers".
	ApplicationVndAppleNumbers = "application/vnd.apple.numbers"

	// ApplicationVndApplePages is the constant for the mime type "application/vnd.apple.pages".
	ApplicationVndApplePages = "application/vnd.apple.pages"

	// ApplicationVndArastraSwi is the constant for the mime type "application/vnd.arastra.swi - OBSOLETED in favor of application/vnd.aristanetworks.swi".
	ApplicationVndArastraSwi = "application/vnd.arastra.swi"

	// ApplicationVndAristanetworksSwi is the constant for the mime type "application/vnd.aristanetworks.swi".
	ApplicationVndAristanetworksSwi = "application/vnd.aristanetworks.swi"

	// ApplicationVndArtisanJson is the constant for the mime type "application/vnd.artisan+json".
	ApplicationVndArtisanJson = "application/vnd.artisan+json"

	// ApplicationVndArtsquare is the constant for the mime type "application/vnd.artsquare".
	ApplicationVndArtsquare = "application/vnd.artsquare"

	// ApplicationVndAstraeaSoftwareIota is the constant for the mime type "application/vnd.astraea-software.iota".
	ApplicationVndAstraeaSoftwareIota = "application/vnd.astraea-software.iota"

	// ApplicationVndAudiograph is the constant for the mime type "application/vnd.audiograph".
	ApplicationVndAudiograph = "application/vnd.audiograph"

	// ApplicationVndAutopackage is the constant for the mime type "application/vnd.autopackage".
	ApplicationVndAutopackage = "application/vnd.autopackage"

	// ApplicationVndAvalonJson is the constant for the mime type "application/vnd.avalon+json".
	ApplicationVndAvalonJson = "application/vnd.avalon+json"

	// ApplicationVndAvistarXml is the constant for the mime type "application/vnd.avistar+xml".
	ApplicationVndAvistarXml = "application/vnd.avistar+xml"

	// ApplicationVndBalsamiqBmmlXml is the constant for the mime type "application/vnd.balsamiq.bmml+xml".
	ApplicationVndBalsamiqBmmlXml = "application/vnd.balsamiq.bmml+xml"

	// ApplicationVndBananaAccounting is the constant for the mime type "application/vnd.banana-accounting".
	ApplicationVndBananaAccounting = "application/vnd.banana-accounting"

	// ApplicationVndBbfUspError is the constant for the mime type "application/vnd.bbf.usp.error".
	ApplicationVndBbfUspError = "application/vnd.bbf.usp.error"

	// ApplicationVndBbfUspMsg is the constant for the mime type "application/vnd.bbf.usp.msg".
	ApplicationVndBbfUspMsg = "application/vnd.bbf.usp.msg"

	// ApplicationVndBbfUspMsgJson is the constant for the mime type "application/vnd.bbf.usp.msg+json".
	ApplicationVndBbfUspMsgJson = "application/vnd.bbf.usp.msg+json"

	// ApplicationVndBalsamiqBmpr is the constant for the mime type "application/vnd.balsamiq.bmpr".
	ApplicationVndBalsamiqBmpr = "application/vnd.balsamiq.bmpr"

	// ApplicationVndBekitzurStechJson is the constant for the mime type "application/vnd.bekitzur-stech+json".
	ApplicationVndBekitzurStechJson = "application/vnd.bekitzur-stech+json"

	// ApplicationVndBintMedContent is the constant for the mime type "application/vnd.bint.med-content".
	ApplicationVndBintMedContent = "application/vnd.bint.med-content"

	// ApplicationVndBiopaxRdfXml is the constant for the mime type "application/vnd.biopax.rdf+xml".
	ApplicationVndBiopaxRdfXml = "application/vnd.biopax.rdf+xml"

	// ApplicationVndBlinkIdbValueWrapper is the constant for the mime type "application/vnd.blink-idb-value-wrapper".
	ApplicationVndBlinkIdbValueWrapper = "application/vnd.blink-idb-value-wrapper"

	// ApplicationVndBlueiceMultipass is the constant for the mime type "application/vnd.blueice.multipass".
	ApplicationVndBlueiceMultipass = "application/vnd.blueice.multipass"

	// ApplicationVndBluetoothEpOob is the constant for the mime type "application/vnd.bluetooth.ep.oob".
	ApplicationVndBluetoothEpOob = "application/vnd.bluetooth.ep.oob"

	// ApplicationVndBluetoothLeOob is the constant for the mime type "application/vnd.bluetooth.le.oob".
	ApplicationVndBluetoothLeOob = "application/vnd.bluetooth.le.oob"

	// ApplicationVndBmi is the constant for the mime type "application/vnd.bmi".
	ApplicationVndBmi = "application/vnd.bmi"

	// ApplicationVndBpf is the constant for the mime type "application/vnd.bpf".
	ApplicationVndBpf = "application/vnd.bpf"

	// ApplicationVndBpf3 is the constant for the mime type "application/vnd.bpf3".
	ApplicationVndBpf3 = "application/vnd.bpf3"

	// ApplicationVndBusinessobjects is the constant for the mime type "application/vnd.businessobjects".
	ApplicationVndBusinessobjects = "application/vnd.businessobjects"

	// ApplicationVndByuUapiJson is the constant for the mime type "application/vnd.byu.uapi+json".
	ApplicationVndByuUapiJson = "application/vnd.byu.uapi+json"

	// ApplicationVndCabJscript is the constant for the mime type "application/vnd.cab-jscript".
	ApplicationVndCabJscript = "application/vnd.cab-jscript"

	// ApplicationVndCanonCpdl is the constant for the mime type "application/vnd.canon-cpdl".
	ApplicationVndCanonCpdl = "application/vnd.canon-cpdl"

	// ApplicationVndCanonLips is the constant for the mime type "application/vnd.canon-lips".
	ApplicationVndCanonLips = "application/vnd.canon-lips"

	// ApplicationVndCapasystemsPgJson is the constant for the mime type "application/vnd.capasystems-pg+json".
	ApplicationVndCapasystemsPgJson = "application/vnd.capasystems-pg+json"

	// ApplicationVndCendioThinlincClientconf is the constant for the mime type "application/vnd.cendio.thinlinc.clientconf".
	ApplicationVndCendioThinlincClientconf = "application/vnd.cendio.thinlinc.clientconf"

	// ApplicationVndCenturySystemsTcpStream is the constant for the mime type "application/vnd.century-systems.tcp_stream".
	ApplicationVndCenturySystemsTcpStream = "application/vnd.century-systems.tcp_stream"

	// ApplicationVndChemdrawXml is the constant for the mime type "application/vnd.chemdraw+xml".
	ApplicationVndChemdrawXml = "application/vnd.chemdraw+xml"

	// ApplicationVndChessPgn is the constant for the mime type "application/vnd.chess-pgn".
	ApplicationVndChessPgn = "application/vnd.chess-pgn"

	// ApplicationVndChipnutsKaraokeMmd is the constant for the mime type "application/vnd.chipnuts.karaoke-mmd".
	ApplicationVndChipnutsKaraokeMmd = "application/vnd.chipnuts.karaoke-mmd"

	// ApplicationVndCiedi is the constant for the mime type "application/vnd.ciedi".
	ApplicationVndCiedi = "application/vnd.ciedi"

	// ApplicationVndCinderella is the constant for the mime type "application/vnd.cinderella".
	ApplicationVndCinderella = "application/vnd.cinderella"

	// ApplicationVndCirpackIsdnExt is the constant for the mime type "application/vnd.cirpack.isdn-ext".
	ApplicationVndCirpackIsdnExt = "application/vnd.cirpack.isdn-ext"

	// ApplicationVndCitationstylesStyleXml is the constant for the mime type "application/vnd.citationstyles.style+xml".
	ApplicationVndCitationstylesStyleXml = "application/vnd.citationstyles.style+xml"

	// ApplicationVndClaymore is the constant for the mime type "application/vnd.claymore".
	ApplicationVndClaymore = "application/vnd.claymore"

	// ApplicationVndCloantoRp9 is the constant for the mime type "application/vnd.cloanto.rp9".
	ApplicationVndCloantoRp9 = "application/vnd.cloanto.rp9"

	// ApplicationVndClonkC4Group is the constant for the mime type "application/vnd.clonk.c4group".
	ApplicationVndClonkC4Group = "application/vnd.clonk.c4group"

	// ApplicationVndCluetrustCartomobileConfig is the constant for the mime type "application/vnd.cluetrust.cartomobile-config".
	ApplicationVndCluetrustCartomobileConfig = "application/vnd.cluetrust.cartomobile-config"

	// ApplicationVndCluetrustCartomobileConfigPkg is the constant for the mime type "application/vnd.cluetrust.cartomobile-config-pkg".
	ApplicationVndCluetrustCartomobileConfigPkg = "application/vnd.cluetrust.cartomobile-config-pkg"

	// ApplicationVndCoffeescript is the constant for the mime type "application/vnd.coffeescript".
	ApplicationVndCoffeescript = "application/vnd.coffeescript"

	// ApplicationVndCollabioXodocumentsDocument is the constant for the mime type "application/vnd.collabio.xodocuments.document".
	ApplicationVndCollabioXodocumentsDocument = "application/vnd.collabio.xodocuments.document"

	// ApplicationVndCollabioXodocumentsDocumentTemplate is the constant for the mime type "application/vnd.collabio.xodocuments.document-template".
	ApplicationVndCollabioXodocumentsDocumentTemplate = "application/vnd.collabio.xodocuments.document-template"

	// ApplicationVndCollabioXodocumentsPresentation is the constant for the mime type "application/vnd.collabio.xodocuments.presentation".
	ApplicationVndCollabioXodocumentsPresentation = "application/vnd.collabio.xodocuments.presentation"

	// ApplicationVndCollabioXodocumentsPresentationTemplate is the constant for the mime type "application/vnd.collabio.xodocuments.presentation-template".
	ApplicationVndCollabioXodocumentsPresentationTemplate = "application/vnd.collabio.xodocuments.presentation-template"

	// ApplicationVndCollabioXodocumentsSpreadsheet is the constant for the mime type "application/vnd.collabio.xodocuments.spreadsheet".
	ApplicationVndCollabioXodocumentsSpreadsheet = "application/vnd.collabio.xodocuments.spreadsheet"

	// ApplicationVndCollabioXodocumentsSpreadsheetTemplate is the constant for the mime type "application/vnd.collabio.xodocuments.spreadsheet-template".
	ApplicationVndCollabioXodocumentsSpreadsheetTemplate = "application/vnd.collabio.xodocuments.spreadsheet-template"

	// ApplicationVndCollectionDocJson is the constant for the mime type "application/vnd.collection.doc+json".
	ApplicationVndCollectionDocJson = "application/vnd.collection.doc+json"

	// ApplicationVndCollectionJson is the constant for the mime type "application/vnd.collection+json".
	ApplicationVndCollectionJson = "application/vnd.collection+json"

	// ApplicationVndCollectionNextJson is the constant for the mime type "application/vnd.collection.next+json".
	ApplicationVndCollectionNextJson = "application/vnd.collection.next+json"

	// ApplicationVndComicbookRar is the constant for the mime type "application/vnd.comicbook-rar".
	ApplicationVndComicbookRar = "application/vnd.comicbook-rar"

	// ApplicationVndComicbookZip is the constant for the mime type "application/vnd.comicbook+zip".
	ApplicationVndComicbookZip = "application/vnd.comicbook+zip"

	// ApplicationVndCommerceBattelle is the constant for the mime type "application/vnd.commerce-battelle".
	ApplicationVndCommerceBattelle = "application/vnd.commerce-battelle"

	// ApplicationVndCommonspace is the constant for the mime type "application/vnd.commonspace".
	ApplicationVndCommonspace = "application/vnd.commonspace"

	// ApplicationVndCoreosIgnitionJson is the constant for the mime type "application/vnd.coreos.ignition+json".
	ApplicationVndCoreosIgnitionJson = "application/vnd.coreos.ignition+json"

	// ApplicationVndCosmocaller is the constant for the mime type "application/vnd.cosmocaller".
	ApplicationVndCosmocaller = "application/vnd.cosmocaller"

	// ApplicationVndContactCmsg is the constant for the mime type "application/vnd.contact.cmsg".
	ApplicationVndContactCmsg = "application/vnd.contact.cmsg"

	// ApplicationVndCrickClicker is the constant for the mime type "application/vnd.crick.clicker".
	ApplicationVndCrickClicker = "application/vnd.crick.clicker"

	// ApplicationVndCrickClickerKeyboard is the constant for the mime type "application/vnd.crick.clicker.keyboard".
	ApplicationVndCrickClickerKeyboard = "application/vnd.crick.clicker.keyboard"

	// ApplicationVndCrickClickerPalette is the constant for the mime type "application/vnd.crick.clicker.palette".
	ApplicationVndCrickClickerPalette = "application/vnd.crick.clicker.palette"

	// ApplicationVndCrickClickerTemplate is the constant for the mime type "application/vnd.crick.clicker.template".
	ApplicationVndCrickClickerTemplate = "application/vnd.crick.clicker.template"

	// ApplicationVndCrickClickerWordbank is the constant for the mime type "application/vnd.crick.clicker.wordbank".
	ApplicationVndCrickClickerWordbank = "application/vnd.crick.clicker.wordbank"

	// ApplicationVndCriticaltoolsWbsXml is the constant for the mime type "application/vnd.criticaltools.wbs+xml".
	ApplicationVndCriticaltoolsWbsXml = "application/vnd.criticaltools.wbs+xml"

	// ApplicationVndCryptiiPipeJson is the constant for the mime type "application/vnd.cryptii.pipe+json".
	ApplicationVndCryptiiPipeJson = "application/vnd.cryptii.pipe+json"

	// ApplicationVndCryptoShadeFile is the constant for the mime type "application/vnd.crypto-shade-file".
	ApplicationVndCryptoShadeFile = "application/vnd.crypto-shade-file"

	// ApplicationVndCtcPosml is the constant for the mime type "application/vnd.ctc-posml".
	ApplicationVndCtcPosml = "application/vnd.ctc-posml"

	// ApplicationVndCtctWsXml is the constant for the mime type "application/vnd.ctct.ws+xml".
	ApplicationVndCtctWsXml = "application/vnd.ctct.ws+xml"

	// ApplicationVndCupsPdf is the constant for the mime type "application/vnd.cups-pdf".
	ApplicationVndCupsPdf = "application/vnd.cups-pdf"

	// ApplicationVndCupsPostscript is the constant for the mime type "application/vnd.cups-postscript".
	ApplicationVndCupsPostscript = "application/vnd.cups-postscript"

	// ApplicationVndCupsPpd is the constant for the mime type "application/vnd.cups-ppd".
	ApplicationVndCupsPpd = "application/vnd.cups-ppd"

	// ApplicationVndCupsRaster is the constant for the mime type "application/vnd.cups-raster".
	ApplicationVndCupsRaster = "application/vnd.cups-raster"

	// ApplicationVndCupsRaw is the constant for the mime type "application/vnd.cups-raw".
	ApplicationVndCupsRaw = "application/vnd.cups-raw"

	// ApplicationVndCurl is the constant for the mime type "application/vnd.curl".
	ApplicationVndCurl = "application/vnd.curl"

	// ApplicationVndCyanDeanRootXml is the constant for the mime type "application/vnd.cyan.dean.root+xml".
	ApplicationVndCyanDeanRootXml = "application/vnd.cyan.dean.root+xml"

	// ApplicationVndCybank is the constant for the mime type "application/vnd.cybank".
	ApplicationVndCybank = "application/vnd.cybank"

	// ApplicationVndD2LCoursepackage1P0Zip is the constant for the mime type "application/vnd.d2l.coursepackage1p0+zip".
	ApplicationVndD2LCoursepackage1P0Zip = "application/vnd.d2l.coursepackage1p0+zip"

	// ApplicationVndDart is the constant for the mime type "application/vnd.dart".
	ApplicationVndDart = "application/vnd.dart"

	// ApplicationVndDataVisionRdz is the constant for the mime type "application/vnd.data-vision.rdz".
	ApplicationVndDataVisionRdz = "application/vnd.data-vision.rdz"

	// ApplicationVndDatapackageJson is the constant for the mime type "application/vnd.datapackage+json".
	ApplicationVndDatapackageJson = "application/vnd.datapackage+json"

	// ApplicationVndDataresourceJson is the constant for the mime type "application/vnd.dataresource+json".
	ApplicationVndDataresourceJson = "application/vnd.dataresource+json"

	// ApplicationVndDebianBinaryPackage is the constant for the mime type "application/vnd.debian.binary-package".
	ApplicationVndDebianBinaryPackage = "application/vnd.debian.binary-package"

	// ApplicationVndDeceData is the constant for the mime type "application/vnd.dece.data".
	ApplicationVndDeceData = "application/vnd.dece.data"

	// ApplicationVndDeceTtmlXml is the constant for the mime type "application/vnd.dece.ttml+xml".
	ApplicationVndDeceTtmlXml = "application/vnd.dece.ttml+xml"

	// ApplicationVndDeceUnspecified is the constant for the mime type "application/vnd.dece.unspecified".
	ApplicationVndDeceUnspecified = "application/vnd.dece.unspecified"

	// ApplicationVndDeceZip is the constant for the mime type "application/vnd.dece.zip".
	ApplicationVndDeceZip = "application/vnd.dece.zip"

	// ApplicationVndDenovoFcselayoutLink is the constant for the mime type "application/vnd.denovo.fcselayout-link".
	ApplicationVndDenovoFcselayoutLink = "application/vnd.denovo.fcselayout-link"

	// ApplicationVndDesmumeMovie is the constant for the mime type "application/vnd.desmume.movie".
	ApplicationVndDesmumeMovie = "application/vnd.desmume.movie"

	// ApplicationVndDirBiPlateDlNosuffix is the constant for the mime type "application/vnd.dir-bi.plate-dl-nosuffix".
	ApplicationVndDirBiPlateDlNosuffix = "application/vnd.dir-bi.plate-dl-nosuffix"

	// ApplicationVndDmDelegationXml is the constant for the mime type "application/vnd.dm.delegation+xml".
	ApplicationVndDmDelegationXml = "application/vnd.dm.delegation+xml"

	// ApplicationVndDna is the constant for the mime type "application/vnd.dna".
	ApplicationVndDna = "application/vnd.dna"

	// ApplicationVndDocumentJson is the constant for the mime type "application/vnd.document+json".
	ApplicationVndDocumentJson = "application/vnd.document+json"

	// ApplicationVndDolbyMobile1 is the constant for the mime type "application/vnd.dolby.mobile.1".
	ApplicationVndDolbyMobile1 = "application/vnd.dolby.mobile.1"

	// ApplicationVndDolbyMobile2 is the constant for the mime type "application/vnd.dolby.mobile.2".
	ApplicationVndDolbyMobile2 = "application/vnd.dolby.mobile.2"

	// ApplicationVndDoremirScorecloudBinaryDocument is the constant for the mime type "application/vnd.doremir.scorecloud-binary-document".
	ApplicationVndDoremirScorecloudBinaryDocument = "application/vnd.doremir.scorecloud-binary-document"

	// ApplicationVndDpgraph is the constant for the mime type "application/vnd.dpgraph".
	ApplicationVndDpgraph = "application/vnd.dpgraph"

	// ApplicationVndDreamfactory is the constant for the mime type "application/vnd.dreamfactory".
	ApplicationVndDreamfactory = "application/vnd.dreamfactory"

	// ApplicationVndDriveJson is the constant for the mime type "application/vnd.drive+json".
	ApplicationVndDriveJson = "application/vnd.drive+json"

	// ApplicationVndDtgLocal is the constant for the mime type "application/vnd.dtg.local".
	ApplicationVndDtgLocal = "application/vnd.dtg.local"

	// ApplicationVndDtgLocalFlash is the constant for the mime type "application/vnd.dtg.local.flash".
	ApplicationVndDtgLocalFlash = "application/vnd.dtg.local.flash"

	// ApplicationVndDtgLocalHtml is the constant for the mime type "application/vnd.dtg.local.html".
	ApplicationVndDtgLocalHtml = "application/vnd.dtg.local.html"

	// ApplicationVndDvbAit is the constant for the mime type "application/vnd.dvb.ait".
	ApplicationVndDvbAit = "application/vnd.dvb.ait"

	// ApplicationVndDvbDvbj is the constant for the mime type "application/vnd.dvb.dvbj".
	ApplicationVndDvbDvbj = "application/vnd.dvb.dvbj"

	// ApplicationVndDvbEsgcontainer is the constant for the mime type "application/vnd.dvb.esgcontainer".
	ApplicationVndDvbEsgcontainer = "application/vnd.dvb.esgcontainer"

	// ApplicationVndDvbIpdcdftnotifaccess is the constant for the mime type "application/vnd.dvb.ipdcdftnotifaccess".
	ApplicationVndDvbIpdcdftnotifaccess = "application/vnd.dvb.ipdcdftnotifaccess"

	// ApplicationVndDvbIpdcesgaccess is the constant for the mime type "application/vnd.dvb.ipdcesgaccess".
	ApplicationVndDvbIpdcesgaccess = "application/vnd.dvb.ipdcesgaccess"

	// ApplicationVndDvbIpdcesgaccess2 is the constant for the mime type "application/vnd.dvb.ipdcesgaccess2".
	ApplicationVndDvbIpdcesgaccess2 = "application/vnd.dvb.ipdcesgaccess2"

	// ApplicationVndDvbIpdcesgpdd is the constant for the mime type "application/vnd.dvb.ipdcesgpdd".
	ApplicationVndDvbIpdcesgpdd = "application/vnd.dvb.ipdcesgpdd"

	// ApplicationVndDvbIpdcroaming is the constant for the mime type "application/vnd.dvb.ipdcroaming".
	ApplicationVndDvbIpdcroaming = "application/vnd.dvb.ipdcroaming"

	// ApplicationVndDvbIptvAlfecBase is the constant for the mime type "application/vnd.dvb.iptv.alfec-base".
	ApplicationVndDvbIptvAlfecBase = "application/vnd.dvb.iptv.alfec-base"

	// ApplicationVndDvbIptvAlfecEnhancement is the constant for the mime type "application/vnd.dvb.iptv.alfec-enhancement".
	ApplicationVndDvbIptvAlfecEnhancement = "application/vnd.dvb.iptv.alfec-enhancement"

	// ApplicationVndDvbNotifAggregateRootXml is the constant for the mime type "application/vnd.dvb.notif-aggregate-root+xml".
	ApplicationVndDvbNotifAggregateRootXml = "application/vnd.dvb.notif-aggregate-root+xml"

	// ApplicationVndDvbNotifContainerXml is the constant for the mime type "application/vnd.dvb.notif-container+xml".
	ApplicationVndDvbNotifContainerXml = "application/vnd.dvb.notif-container+xml"

	// ApplicationVndDvbNotifGenericXml is the constant for the mime type "application/vnd.dvb.notif-generic+xml".
	ApplicationVndDvbNotifGenericXml = "application/vnd.dvb.notif-generic+xml"

	// ApplicationVndDvbNotifIaMsglistXml is the constant for the mime type "application/vnd.dvb.notif-ia-msglist+xml".
	ApplicationVndDvbNotifIaMsglistXml = "application/vnd.dvb.notif-ia-msglist+xml"

	// ApplicationVndDvbNotifIaRegistrationRequestXml is the constant for the mime type "application/vnd.dvb.notif-ia-registration-request+xml".
	ApplicationVndDvbNotifIaRegistrationRequestXml = "application/vnd.dvb.notif-ia-registration-request+xml"

	// ApplicationVndDvbNotifIaRegistrationResponseXml is the constant for the mime type "application/vnd.dvb.notif-ia-registration-response+xml".
	ApplicationVndDvbNotifIaRegistrationResponseXml = "application/vnd.dvb.notif-ia-registration-response+xml"

	// ApplicationVndDvbNotifInitXml is the constant for the mime type "application/vnd.dvb.notif-init+xml".
	ApplicationVndDvbNotifInitXml = "application/vnd.dvb.notif-init+xml"

	// ApplicationVndDvbPfr is the constant for the mime type "application/vnd.dvb.pfr".
	ApplicationVndDvbPfr = "application/vnd.dvb.pfr"

	// ApplicationVndDvbService is the constant for the mime type "application/vnd.dvb.service".
	ApplicationVndDvbService = "application/vnd.dvb.service"

	// ApplicationVndDxr is the constant for the mime type "application/vnd.dxr".
	ApplicationVndDxr = "application/vnd.dxr"

	// ApplicationVndDynageo is the constant for the mime type "application/vnd.dynageo".
	ApplicationVndDynageo = "application/vnd.dynageo"

	// ApplicationVndDzr is the constant for the mime type "application/vnd.dzr".
	ApplicationVndDzr = "application/vnd.dzr"

	// ApplicationVndEasykaraokeCdgdownload is the constant for the mime type "application/vnd.easykaraoke.cdgdownload".
	ApplicationVndEasykaraokeCdgdownload = "application/vnd.easykaraoke.cdgdownload"

	// ApplicationVndEcipRlp is the constant for the mime type "application/vnd.ecip.rlp".
	ApplicationVndEcipRlp = "application/vnd.ecip.rlp"

	// ApplicationVndEcdisUpdate is the constant for the mime type "application/vnd.ecdis-update".
	ApplicationVndEcdisUpdate = "application/vnd.ecdis-update"

	// ApplicationVndEcowinChart is the constant for the mime type "application/vnd.ecowin.chart".
	ApplicationVndEcowinChart = "application/vnd.ecowin.chart"

	// ApplicationVndEcowinFilerequest is the constant for the mime type "application/vnd.ecowin.filerequest".
	ApplicationVndEcowinFilerequest = "application/vnd.ecowin.filerequest"

	// ApplicationVndEcowinFileupdate is the constant for the mime type "application/vnd.ecowin.fileupdate".
	ApplicationVndEcowinFileupdate = "application/vnd.ecowin.fileupdate"

	// ApplicationVndEcowinSeries is the constant for the mime type "application/vnd.ecowin.series".
	ApplicationVndEcowinSeries = "application/vnd.ecowin.series"

	// ApplicationVndEcowinSeriesrequest is the constant for the mime type "application/vnd.ecowin.seriesrequest".
	ApplicationVndEcowinSeriesrequest = "application/vnd.ecowin.seriesrequest"

	// ApplicationVndEcowinSeriesupdate is the constant for the mime type "application/vnd.ecowin.seriesupdate".
	ApplicationVndEcowinSeriesupdate = "application/vnd.ecowin.seriesupdate"

	// ApplicationVndEfiImg is the constant for the mime type "application/vnd.efi.img".
	ApplicationVndEfiImg = "application/vnd.efi.img"

	// ApplicationVndEfiIso is the constant for the mime type "application/vnd.efi.iso".
	ApplicationVndEfiIso = "application/vnd.efi.iso"

	// ApplicationVndEmclientAccessrequestXml is the constant for the mime type "application/vnd.emclient.accessrequest+xml".
	ApplicationVndEmclientAccessrequestXml = "application/vnd.emclient.accessrequest+xml"

	// ApplicationVndEnliven is the constant for the mime type "application/vnd.enliven".
	ApplicationVndEnliven = "application/vnd.enliven"

	// ApplicationVndEnphaseEnvoy is the constant for the mime type "application/vnd.enphase.envoy".
	ApplicationVndEnphaseEnvoy = "application/vnd.enphase.envoy"

	// ApplicationVndEprintsDataXml is the constant for the mime type "application/vnd.eprints.data+xml".
	ApplicationVndEprintsDataXml = "application/vnd.eprints.data+xml"

	// ApplicationVndEpsonEsf is the constant for the mime type "application/vnd.epson.esf".
	ApplicationVndEpsonEsf = "application/vnd.epson.esf"

	// ApplicationVndEpsonMsf is the constant for the mime type "application/vnd.epson.msf".
	ApplicationVndEpsonMsf = "application/vnd.epson.msf"

	// ApplicationVndEpsonQuickanime is the constant for the mime type "application/vnd.epson.quickanime".
	ApplicationVndEpsonQuickanime = "application/vnd.epson.quickanime"

	// ApplicationVndEpsonSalt is the constant for the mime type "application/vnd.epson.salt".
	ApplicationVndEpsonSalt = "application/vnd.epson.salt"

	// ApplicationVndEpsonSsf is the constant for the mime type "application/vnd.epson.ssf".
	ApplicationVndEpsonSsf = "application/vnd.epson.ssf"

	// ApplicationVndEricssonQuickcall is the constant for the mime type "application/vnd.ericsson.quickcall".
	ApplicationVndEricssonQuickcall = "application/vnd.ericsson.quickcall"

	// ApplicationVndEspassEspassZip is the constant for the mime type "application/vnd.espass-espass+zip".
	ApplicationVndEspassEspassZip = "application/vnd.espass-espass+zip"

	// ApplicationVndEszigno3Xml is the constant for the mime type "application/vnd.eszigno3+xml".
	ApplicationVndEszigno3Xml = "application/vnd.eszigno3+xml"

	// ApplicationVndEtsiAocXml is the constant for the mime type "application/vnd.etsi.aoc+xml".
	ApplicationVndEtsiAocXml = "application/vnd.etsi.aoc+xml"

	// ApplicationVndEtsiAsicSZip is the constant for the mime type "application/vnd.etsi.asic-s+zip".
	ApplicationVndEtsiAsicSZip = "application/vnd.etsi.asic-s+zip"

	// ApplicationVndEtsiAsicEZip is the constant for the mime type "application/vnd.etsi.asic-e+zip".
	ApplicationVndEtsiAsicEZip = "application/vnd.etsi.asic-e+zip"

	// ApplicationVndEtsiCugXml is the constant for the mime type "application/vnd.etsi.cug+xml".
	ApplicationVndEtsiCugXml = "application/vnd.etsi.cug+xml"

	// ApplicationVndEtsiIptvcommandXml is the constant for the mime type "application/vnd.etsi.iptvcommand+xml".
	ApplicationVndEtsiIptvcommandXml = "application/vnd.etsi.iptvcommand+xml"

	// ApplicationVndEtsiIptvdiscoveryXml is the constant for the mime type "application/vnd.etsi.iptvdiscovery+xml".
	ApplicationVndEtsiIptvdiscoveryXml = "application/vnd.etsi.iptvdiscovery+xml"

	// ApplicationVndEtsiIptvprofileXml is the constant for the mime type "application/vnd.etsi.iptvprofile+xml".
	ApplicationVndEtsiIptvprofileXml = "application/vnd.etsi.iptvprofile+xml"

	// ApplicationVndEtsiIptvsadBcXml is the constant for the mime type "application/vnd.etsi.iptvsad-bc+xml".
	ApplicationVndEtsiIptvsadBcXml = "application/vnd.etsi.iptvsad-bc+xml"

	// ApplicationVndEtsiIptvsadCodXml is the constant for the mime type "application/vnd.etsi.iptvsad-cod+xml".
	ApplicationVndEtsiIptvsadCodXml = "application/vnd.etsi.iptvsad-cod+xml"

	// ApplicationVndEtsiIptvsadNpvrXml is the constant for the mime type "application/vnd.etsi.iptvsad-npvr+xml".
	ApplicationVndEtsiIptvsadNpvrXml = "application/vnd.etsi.iptvsad-npvr+xml"

	// ApplicationVndEtsiIptvserviceXml is the constant for the mime type "application/vnd.etsi.iptvservice+xml".
	ApplicationVndEtsiIptvserviceXml = "application/vnd.etsi.iptvservice+xml"

	// ApplicationVndEtsiIptvsyncXml is the constant for the mime type "application/vnd.etsi.iptvsync+xml".
	ApplicationVndEtsiIptvsyncXml = "application/vnd.etsi.iptvsync+xml"

	// ApplicationVndEtsiIptvueprofileXml is the constant for the mime type "application/vnd.etsi.iptvueprofile+xml".
	ApplicationVndEtsiIptvueprofileXml = "application/vnd.etsi.iptvueprofile+xml"

	// ApplicationVndEtsiMcidXml is the constant for the mime type "application/vnd.etsi.mcid+xml".
	ApplicationVndEtsiMcidXml = "application/vnd.etsi.mcid+xml"

	// ApplicationVndEtsiMheg5 is the constant for the mime type "application/vnd.etsi.mheg5".
	ApplicationVndEtsiMheg5 = "application/vnd.etsi.mheg5"

	// ApplicationVndEtsiOverloadControlPolicyDatasetXml is the constant for the mime type "application/vnd.etsi.overload-control-policy-dataset+xml".
	ApplicationVndEtsiOverloadControlPolicyDatasetXml = "application/vnd.etsi.overload-control-policy-dataset+xml"

	// ApplicationVndEtsiPstnXml is the constant for the mime type "application/vnd.etsi.pstn+xml".
	ApplicationVndEtsiPstnXml = "application/vnd.etsi.pstn+xml"

	// ApplicationVndEtsiSciXml is the constant for the mime type "application/vnd.etsi.sci+xml".
	ApplicationVndEtsiSciXml = "application/vnd.etsi.sci+xml"

	// ApplicationVndEtsiSimservsXml is the constant for the mime type "application/vnd.etsi.simservs+xml".
	ApplicationVndEtsiSimservsXml = "application/vnd.etsi.simservs+xml"

	// ApplicationVndEtsiTimestampToken is the constant for the mime type "application/vnd.etsi.timestamp-token".
	ApplicationVndEtsiTimestampToken = "application/vnd.etsi.timestamp-token"

	// ApplicationVndEtsiTslXml is the constant for the mime type "application/vnd.etsi.tsl+xml".
	ApplicationVndEtsiTslXml = "application/vnd.etsi.tsl+xml"

	// ApplicationVndEtsiTslDer is the constant for the mime type "application/vnd.etsi.tsl.der".
	ApplicationVndEtsiTslDer = "application/vnd.etsi.tsl.der"

	// ApplicationVndEvolvEcigProfile is the constant for the mime type "application/vnd.evolv.ecig.profile".
	ApplicationVndEvolvEcigProfile = "application/vnd.evolv.ecig.profile"

	// ApplicationVndEvolvEcigSettings is the constant for the mime type "application/vnd.evolv.ecig.settings".
	ApplicationVndEvolvEcigSettings = "application/vnd.evolv.ecig.settings"

	// ApplicationVndEvolvEcigTheme is the constant for the mime type "application/vnd.evolv.ecig.theme".
	ApplicationVndEvolvEcigTheme = "application/vnd.evolv.ecig.theme"

	// ApplicationVndEudoraData is the constant for the mime type "application/vnd.eudora.data".
	ApplicationVndEudoraData = "application/vnd.eudora.data"

	// ApplicationVndExstreamEmpowerZip is the constant for the mime type "application/vnd.exstream-empower+zip".
	ApplicationVndExstreamEmpowerZip = "application/vnd.exstream-empower+zip"

	// ApplicationVndExstreamPackage is the constant for the mime type "application/vnd.exstream-package".
	ApplicationVndExstreamPackage = "application/vnd.exstream-package"

	// ApplicationVndEzpixAlbum is the constant for the mime type "application/vnd.ezpix-album".
	ApplicationVndEzpixAlbum = "application/vnd.ezpix-album"

	// ApplicationVndEzpixPackage is the constant for the mime type "application/vnd.ezpix-package".
	ApplicationVndEzpixPackage = "application/vnd.ezpix-package"

	// ApplicationVndFSecureMobile is the constant for the mime type "application/vnd.f-secure.mobile".
	ApplicationVndFSecureMobile = "application/vnd.f-secure.mobile"

	// ApplicationVndFastcopyDiskImage is the constant for the mime type "application/vnd.fastcopy-disk-image".
	ApplicationVndFastcopyDiskImage = "application/vnd.fastcopy-disk-image"

	// ApplicationVndFdf is the constant for the mime type "application/vnd.fdf".
	ApplicationVndFdf = "application/vnd.fdf"

	// ApplicationVndFdsnMseed is the constant for the mime type "application/vnd.fdsn.mseed".
	ApplicationVndFdsnMseed = "application/vnd.fdsn.mseed"

	// ApplicationVndFdsnSeed is the constant for the mime type "application/vnd.fdsn.seed".
	ApplicationVndFdsnSeed = "application/vnd.fdsn.seed"

	// ApplicationVndFfsns is the constant for the mime type "application/vnd.ffsns".
	ApplicationVndFfsns = "application/vnd.ffsns"

	// ApplicationVndFilmitZfc is the constant for the mime type "application/vnd.filmit.zfc".
	ApplicationVndFilmitZfc = "application/vnd.filmit.zfc"

	// ApplicationVndFints is the constant for the mime type "application/vnd.fints".
	ApplicationVndFints = "application/vnd.fints"

	// ApplicationVndFiremonkeysCloudcell is the constant for the mime type "application/vnd.firemonkeys.cloudcell".
	ApplicationVndFiremonkeysCloudcell = "application/vnd.firemonkeys.cloudcell"

	// ApplicationVndFloGraphIt is the constant for the mime type "application/vnd.FloGraphIt".
	ApplicationVndFloGraphIt = "application/vnd.FloGraphIt"

	// ApplicationVndFluxtimeClip is the constant for the mime type "application/vnd.fluxtime.clip".
	ApplicationVndFluxtimeClip = "application/vnd.fluxtime.clip"

	// ApplicationVndFontFontforgeSfd is the constant for the mime type "application/vnd.font-fontforge-sfd".
	ApplicationVndFontFontforgeSfd = "application/vnd.font-fontforge-sfd"

	// ApplicationVndFramemaker is the constant for the mime type "application/vnd.framemaker".
	ApplicationVndFramemaker = "application/vnd.framemaker"

	// ApplicationVndFrogansFnc is the constant for the mime type "application/vnd.frogans.fnc".
	ApplicationVndFrogansFnc = "application/vnd.frogans.fnc"

	// ApplicationVndFrogansLtf is the constant for the mime type "application/vnd.frogans.ltf".
	ApplicationVndFrogansLtf = "application/vnd.frogans.ltf"

	// ApplicationVndFscWeblaunch is the constant for the mime type "application/vnd.fsc.weblaunch".
	ApplicationVndFscWeblaunch = "application/vnd.fsc.weblaunch"

	// ApplicationVndFujitsuOasys is the constant for the mime type "application/vnd.fujitsu.oasys".
	ApplicationVndFujitsuOasys = "application/vnd.fujitsu.oasys"

	// ApplicationVndFujitsuOasys2 is the constant for the mime type "application/vnd.fujitsu.oasys2".
	ApplicationVndFujitsuOasys2 = "application/vnd.fujitsu.oasys2"

	// ApplicationVndFujitsuOasys3 is the constant for the mime type "application/vnd.fujitsu.oasys3".
	ApplicationVndFujitsuOasys3 = "application/vnd.fujitsu.oasys3"

	// ApplicationVndFujitsuOasysgp is the constant for the mime type "application/vnd.fujitsu.oasysgp".
	ApplicationVndFujitsuOasysgp = "application/vnd.fujitsu.oasysgp"

	// ApplicationVndFujitsuOasysprs is the constant for the mime type "application/vnd.fujitsu.oasysprs".
	ApplicationVndFujitsuOasysprs = "application/vnd.fujitsu.oasysprs"

	// ApplicationVndFujixeroxART4 is the constant for the mime type "application/vnd.fujixerox.ART4".
	ApplicationVndFujixeroxART4 = "application/vnd.fujixerox.ART4"

	// ApplicationVndFujixeroxARTEX is the constant for the mime type "application/vnd.fujixerox.ART-EX".
	ApplicationVndFujixeroxARTEX = "application/vnd.fujixerox.ART-EX"

	// ApplicationVndFujixeroxDdd is the constant for the mime type "application/vnd.fujixerox.ddd".
	ApplicationVndFujixeroxDdd = "application/vnd.fujixerox.ddd"

	// ApplicationVndFujixeroxDocuworks is the constant for the mime type "application/vnd.fujixerox.docuworks".
	ApplicationVndFujixeroxDocuworks = "application/vnd.fujixerox.docuworks"

	// ApplicationVndFujixeroxDocuworksBinder is the constant for the mime type "application/vnd.fujixerox.docuworks.binder".
	ApplicationVndFujixeroxDocuworksBinder = "application/vnd.fujixerox.docuworks.binder"

	// ApplicationVndFujixeroxDocuworksContainer is the constant for the mime type "application/vnd.fujixerox.docuworks.container".
	ApplicationVndFujixeroxDocuworksContainer = "application/vnd.fujixerox.docuworks.container"

	// ApplicationVndFujixeroxHBPL is the constant for the mime type "application/vnd.fujixerox.HBPL".
	ApplicationVndFujixeroxHBPL = "application/vnd.fujixerox.HBPL"

	// ApplicationVndFutMisnet is the constant for the mime type "application/vnd.fut-misnet".
	ApplicationVndFutMisnet = "application/vnd.fut-misnet"

	// ApplicationVndFutoinCbor is the constant for the mime type "application/vnd.futoin+cbor".
	ApplicationVndFutoinCbor = "application/vnd.futoin+cbor"

	// ApplicationVndFutoinJson is the constant for the mime type "application/vnd.futoin+json".
	ApplicationVndFutoinJson = "application/vnd.futoin+json"

	// ApplicationVndFuzzysheet is the constant for the mime type "application/vnd.fuzzysheet".
	ApplicationVndFuzzysheet = "application/vnd.fuzzysheet"

	// ApplicationVndGenomatixTuxedo is the constant for the mime type "application/vnd.genomatix.tuxedo".
	ApplicationVndGenomatixTuxedo = "application/vnd.genomatix.tuxedo"

	// ApplicationVndGeoJson is the constant for the mime type "application/vnd.geo+json (OBSOLETED by [RFC7946] in favor of application/geo+json)".
	ApplicationVndGeoJson = "application/vnd.geo+json"

	// ApplicationVndGeocubeXml is the constant for the mime type "application/vnd.geocube+xml - OBSOLETED by request".
	ApplicationVndGeocubeXml = "application/vnd.geocube+xml"

	// ApplicationVndGeogebraFile is the constant for the mime type "application/vnd.geogebra.file".
	ApplicationVndGeogebraFile = "application/vnd.geogebra.file"

	// ApplicationVndGeogebraTool is the constant for the mime type "application/vnd.geogebra.tool".
	ApplicationVndGeogebraTool = "application/vnd.geogebra.tool"

	// ApplicationVndGeometryExplorer is the constant for the mime type "application/vnd.geometry-explorer".
	ApplicationVndGeometryExplorer = "application/vnd.geometry-explorer"

	// ApplicationVndGeonext is the constant for the mime type "application/vnd.geonext".
	ApplicationVndGeonext = "application/vnd.geonext"

	// ApplicationVndGeoplan is the constant for the mime type "application/vnd.geoplan".
	ApplicationVndGeoplan = "application/vnd.geoplan"

	// ApplicationVndGeospace is the constant for the mime type "application/vnd.geospace".
	ApplicationVndGeospace = "application/vnd.geospace"

	// ApplicationVndGerber is the constant for the mime type "application/vnd.gerber".
	ApplicationVndGerber = "application/vnd.gerber"

	// ApplicationVndGlobalplatformCardContentMgt is the constant for the mime type "application/vnd.globalplatform.card-content-mgt".
	ApplicationVndGlobalplatformCardContentMgt = "application/vnd.globalplatform.card-content-mgt"

	// ApplicationVndGlobalplatformCardContentMgtResponse is the constant for the mime type "application/vnd.globalplatform.card-content-mgt-response".
	ApplicationVndGlobalplatformCardContentMgtResponse = "application/vnd.globalplatform.card-content-mgt-response"

	// ApplicationVndGmx is the constant for the mime type "application/vnd.gmx - DEPRECATED".
	ApplicationVndGmx = "application/vnd.gmx"

	// ApplicationVndGoogleEarthKmlXml is the constant for the mime type "application/vnd.google-earth.kml+xml".
	ApplicationVndGoogleEarthKmlXml = "application/vnd.google-earth.kml+xml"

	// ApplicationVndGoogleEarthKmz is the constant for the mime type "application/vnd.google-earth.kmz".
	ApplicationVndGoogleEarthKmz = "application/vnd.google-earth.kmz"

	// ApplicationVndGovSkEFormXml is the constant for the mime type "application/vnd.gov.sk.e-form+xml".
	ApplicationVndGovSkEFormXml = "application/vnd.gov.sk.e-form+xml"

	// ApplicationVndGovSkEFormZip is the constant for the mime type "application/vnd.gov.sk.e-form+zip".
	ApplicationVndGovSkEFormZip = "application/vnd.gov.sk.e-form+zip"

	// ApplicationVndGovSkXmldatacontainerXml is the constant for the mime type "application/vnd.gov.sk.xmldatacontainer+xml".
	ApplicationVndGovSkXmldatacontainerXml = "application/vnd.gov.sk.xmldatacontainer+xml"

	// ApplicationVndGrafeq is the constant for the mime type "application/vnd.grafeq".
	ApplicationVndGrafeq = "application/vnd.grafeq"

	// ApplicationVndGridmp is the constant for the mime type "application/vnd.gridmp".
	ApplicationVndGridmp = "application/vnd.gridmp"

	// ApplicationVndGrooveAccount is the constant for the mime type "application/vnd.groove-account".
	ApplicationVndGrooveAccount = "application/vnd.groove-account"

	// ApplicationVndGrooveHelp is the constant for the mime type "application/vnd.groove-help".
	ApplicationVndGrooveHelp = "application/vnd.groove-help"

	// ApplicationVndGrooveIdentityMessage is the constant for the mime type "application/vnd.groove-identity-message".
	ApplicationVndGrooveIdentityMessage = "application/vnd.groove-identity-message"

	// ApplicationVndGrooveInjector is the constant for the mime type "application/vnd.groove-injector".
	ApplicationVndGrooveInjector = "application/vnd.groove-injector"

	// ApplicationVndGrooveToolMessage is the constant for the mime type "application/vnd.groove-tool-message".
	ApplicationVndGrooveToolMessage = "application/vnd.groove-tool-message"

	// ApplicationVndGrooveToolTemplate is the constant for the mime type "application/vnd.groove-tool-template".
	ApplicationVndGrooveToolTemplate = "application/vnd.groove-tool-template"

	// ApplicationVndGrooveVcard is the constant for the mime type "application/vnd.groove-vcard".
	ApplicationVndGrooveVcard = "application/vnd.groove-vcard"

	// ApplicationVndHalJson is the constant for the mime type "application/vnd.hal+json".
	ApplicationVndHalJson = "application/vnd.hal+json"

	// ApplicationVndHalXml is the constant for the mime type "application/vnd.hal+xml".
	ApplicationVndHalXml = "application/vnd.hal+xml"

	// ApplicationVndHandHeldEntertainmentXml is the constant for the mime type "application/vnd.HandHeld-Entertainment+xml".
	ApplicationVndHandHeldEntertainmentXml = "application/vnd.HandHeld-Entertainment+xml"

	// ApplicationVndHbci is the constant for the mime type "application/vnd.hbci".
	ApplicationVndHbci = "application/vnd.hbci"

	// ApplicationVndHcJson is the constant for the mime type "application/vnd.hc+json".
	ApplicationVndHcJson = "application/vnd.hc+json"

	// ApplicationVndHclBireports is the constant for the mime type "application/vnd.hcl-bireports".
	ApplicationVndHclBireports = "application/vnd.hcl-bireports"

	// ApplicationVndHdt is the constant for the mime type "application/vnd.hdt".
	ApplicationVndHdt = "application/vnd.hdt"

	// ApplicationVndHerokuJson is the constant for the mime type "application/vnd.heroku+json".
	ApplicationVndHerokuJson = "application/vnd.heroku+json"

	// ApplicationVndHheLessonPlayer is the constant for the mime type "application/vnd.hhe.lesson-player".
	ApplicationVndHheLessonPlayer = "application/vnd.hhe.lesson-player"

	// ApplicationVndHpHPGL is the constant for the mime type "application/vnd.hp-HPGL".
	ApplicationVndHpHPGL = "application/vnd.hp-HPGL"

	// ApplicationVndHpHpid is the constant for the mime type "application/vnd.hp-hpid".
	ApplicationVndHpHpid = "application/vnd.hp-hpid"

	// ApplicationVndHpHps is the constant for the mime type "application/vnd.hp-hps".
	ApplicationVndHpHps = "application/vnd.hp-hps"

	// ApplicationVndHpJlyt is the constant for the mime type "application/vnd.hp-jlyt".
	ApplicationVndHpJlyt = "application/vnd.hp-jlyt"

	// ApplicationVndHpPCL is the constant for the mime type "application/vnd.hp-PCL".
	ApplicationVndHpPCL = "application/vnd.hp-PCL"

	// ApplicationVndHpPCLXL is the constant for the mime type "application/vnd.hp-PCLXL".
	ApplicationVndHpPCLXL = "application/vnd.hp-PCLXL"

	// ApplicationVndHttphone is the constant for the mime type "application/vnd.httphone".
	ApplicationVndHttphone = "application/vnd.httphone"

	// ApplicationVndHydrostatixSofData is the constant for the mime type "application/vnd.hydrostatix.sof-data".
	ApplicationVndHydrostatixSofData = "application/vnd.hydrostatix.sof-data"

	// ApplicationVndHyperItemJson is the constant for the mime type "application/vnd.hyper-item+json".
	ApplicationVndHyperItemJson = "application/vnd.hyper-item+json"

	// ApplicationVndHyperJson is the constant for the mime type "application/vnd.hyper+json".
	ApplicationVndHyperJson = "application/vnd.hyper+json"

	// ApplicationVndHyperdriveJson is the constant for the mime type "application/vnd.hyperdrive+json".
	ApplicationVndHyperdriveJson = "application/vnd.hyperdrive+json"

	// ApplicationVndHzn3dCrossword is the constant for the mime type "application/vnd.hzn-3d-crossword".
	ApplicationVndHzn3dCrossword = "application/vnd.hzn-3d-crossword"

	// ApplicationVndIbmAfplinedata is the constant for the mime type "application/vnd.ibm.afplinedata - OBSOLETED in favor of vnd.afpc.afplinedata".
	ApplicationVndIbmAfplinedata = "application/vnd.ibm.afplinedata"

	// ApplicationVndIbmElectronicMedia is the constant for the mime type "application/vnd.ibm.electronic-media".
	ApplicationVndIbmElectronicMedia = "application/vnd.ibm.electronic-media"

	// ApplicationVndIbmMiniPay is the constant for the mime type "application/vnd.ibm.MiniPay".
	ApplicationVndIbmMiniPay = "application/vnd.ibm.MiniPay"

	// ApplicationVndIbmModcap is the constant for the mime type "application/vnd.ibm.modcap - OBSOLETED in favor of application/vnd.afpc.modca".
	ApplicationVndIbmModcap = "application/vnd.ibm.modcap"

	// ApplicationVndIbmRightsManagement is the constant for the mime type "application/vnd.ibm.rights-management".
	ApplicationVndIbmRightsManagement = "application/vnd.ibm.rights-management"

	// ApplicationVndIbmSecureContainer is the constant for the mime type "application/vnd.ibm.secure-container".
	ApplicationVndIbmSecureContainer = "application/vnd.ibm.secure-container"

	// ApplicationVndIccprofile is the constant for the mime type "application/vnd.iccprofile".
	ApplicationVndIccprofile = "application/vnd.iccprofile"

	// ApplicationVndIeee1905 is the constant for the mime type "application/vnd.ieee.1905".
	ApplicationVndIeee1905 = "application/vnd.ieee.1905"

	// ApplicationVndIgloader is the constant for the mime type "application/vnd.igloader".
	ApplicationVndIgloader = "application/vnd.igloader"

	// ApplicationVndImagemeterFolderZip is the constant for the mime type "application/vnd.imagemeter.folder+zip".
	ApplicationVndImagemeterFolderZip = "application/vnd.imagemeter.folder+zip"

	// ApplicationVndImagemeterImageZip is the constant for the mime type "application/vnd.imagemeter.image+zip".
	ApplicationVndImagemeterImageZip = "application/vnd.imagemeter.image+zip"

	// ApplicationVndImmervisionIvp is the constant for the mime type "application/vnd.immervision-ivp".
	ApplicationVndImmervisionIvp = "application/vnd.immervision-ivp"

	// ApplicationVndImmervisionIvu is the constant for the mime type "application/vnd.immervision-ivu".
	ApplicationVndImmervisionIvu = "application/vnd.immervision-ivu"

	// ApplicationVndImsImsccv1P1 is the constant for the mime type "application/vnd.ims.imsccv1p1".
	ApplicationVndImsImsccv1P1 = "application/vnd.ims.imsccv1p1"

	// ApplicationVndImsImsccv1P2 is the constant for the mime type "application/vnd.ims.imsccv1p2".
	ApplicationVndImsImsccv1P2 = "application/vnd.ims.imsccv1p2"

	// ApplicationVndImsImsccv1P3 is the constant for the mime type "application/vnd.ims.imsccv1p3".
	ApplicationVndImsImsccv1P3 = "application/vnd.ims.imsccv1p3"

	// ApplicationVndImsLisV2ResultJson is the constant for the mime type "application/vnd.ims.lis.v2.result+json".
	ApplicationVndImsLisV2ResultJson = "application/vnd.ims.lis.v2.result+json"

	// ApplicationVndImsLtiV2ToolconsumerprofileJson is the constant for the mime type "application/vnd.ims.lti.v2.toolconsumerprofile+json".
	ApplicationVndImsLtiV2ToolconsumerprofileJson = "application/vnd.ims.lti.v2.toolconsumerprofile+json"

	// ApplicationVndImsLtiV2ToolproxyIdJson is the constant for the mime type "application/vnd.ims.lti.v2.toolproxy.id+json".
	ApplicationVndImsLtiV2ToolproxyIdJson = "application/vnd.ims.lti.v2.toolproxy.id+json"

	// ApplicationVndImsLtiV2ToolproxyJson is the constant for the mime type "application/vnd.ims.lti.v2.toolproxy+json".
	ApplicationVndImsLtiV2ToolproxyJson = "application/vnd.ims.lti.v2.toolproxy+json"

	// ApplicationVndImsLtiV2ToolsettingsJson is the constant for the mime type "application/vnd.ims.lti.v2.toolsettings+json".
	ApplicationVndImsLtiV2ToolsettingsJson = "application/vnd.ims.lti.v2.toolsettings+json"

	// ApplicationVndImsLtiV2ToolsettingsSimpleJson is the constant for the mime type "application/vnd.ims.lti.v2.toolsettings.simple+json".
	ApplicationVndImsLtiV2ToolsettingsSimpleJson = "application/vnd.ims.lti.v2.toolsettings.simple+json"

	// ApplicationVndInformedcontrolRmsXml is the constant for the mime type "application/vnd.informedcontrol.rms+xml".
	ApplicationVndInformedcontrolRmsXml = "application/vnd.informedcontrol.rms+xml"

	// ApplicationVndInfotechProject is the constant for the mime type "application/vnd.infotech.project".
	ApplicationVndInfotechProject = "application/vnd.infotech.project"

	// ApplicationVndInfotechProjectXml is the constant for the mime type "application/vnd.infotech.project+xml".
	ApplicationVndInfotechProjectXml = "application/vnd.infotech.project+xml"

	// ApplicationVndInformixVisionary is the constant for the mime type "application/vnd.informix-visionary - OBSOLETED in favor of application/vnd.visionary".
	ApplicationVndInformixVisionary = "application/vnd.informix-visionary"

	// ApplicationVndInnopathWampNotification is the constant for the mime type "application/vnd.innopath.wamp.notification".
	ApplicationVndInnopathWampNotification = "application/vnd.innopath.wamp.notification"

	// ApplicationVndInsorsIgm is the constant for the mime type "application/vnd.insors.igm".
	ApplicationVndInsorsIgm = "application/vnd.insors.igm"

	// ApplicationVndInterconFormnet is the constant for the mime type "application/vnd.intercon.formnet".
	ApplicationVndInterconFormnet = "application/vnd.intercon.formnet"

	// ApplicationVndIntergeo is the constant for the mime type "application/vnd.intergeo".
	ApplicationVndIntergeo = "application/vnd.intergeo"

	// ApplicationVndIntertrustDigibox is the constant for the mime type "application/vnd.intertrust.digibox".
	ApplicationVndIntertrustDigibox = "application/vnd.intertrust.digibox"

	// ApplicationVndIntertrustNncp is the constant for the mime type "application/vnd.intertrust.nncp".
	ApplicationVndIntertrustNncp = "application/vnd.intertrust.nncp"

	// ApplicationVndIntuQbo is the constant for the mime type "application/vnd.intu.qbo".
	ApplicationVndIntuQbo = "application/vnd.intu.qbo"

	// ApplicationVndIntuQfx is the constant for the mime type "application/vnd.intu.qfx".
	ApplicationVndIntuQfx = "application/vnd.intu.qfx"

	// ApplicationVndIptcG2CatalogitemXml is the constant for the mime type "application/vnd.iptc.g2.catalogitem+xml".
	ApplicationVndIptcG2CatalogitemXml = "application/vnd.iptc.g2.catalogitem+xml"

	// ApplicationVndIptcG2ConceptitemXml is the constant for the mime type "application/vnd.iptc.g2.conceptitem+xml".
	ApplicationVndIptcG2ConceptitemXml = "application/vnd.iptc.g2.conceptitem+xml"

	// ApplicationVndIptcG2KnowledgeitemXml is the constant for the mime type "application/vnd.iptc.g2.knowledgeitem+xml".
	ApplicationVndIptcG2KnowledgeitemXml = "application/vnd.iptc.g2.knowledgeitem+xml"

	// ApplicationVndIptcG2NewsitemXml is the constant for the mime type "application/vnd.iptc.g2.newsitem+xml".
	ApplicationVndIptcG2NewsitemXml = "application/vnd.iptc.g2.newsitem+xml"

	// ApplicationVndIptcG2NewsmessageXml is the constant for the mime type "application/vnd.iptc.g2.newsmessage+xml".
	ApplicationVndIptcG2NewsmessageXml = "application/vnd.iptc.g2.newsmessage+xml"

	// ApplicationVndIptcG2PackageitemXml is the constant for the mime type "application/vnd.iptc.g2.packageitem+xml".
	ApplicationVndIptcG2PackageitemXml = "application/vnd.iptc.g2.packageitem+xml"

	// ApplicationVndIptcG2PlanningitemXml is the constant for the mime type "application/vnd.iptc.g2.planningitem+xml".
	ApplicationVndIptcG2PlanningitemXml = "application/vnd.iptc.g2.planningitem+xml"

	// ApplicationVndIpunpluggedRcprofile is the constant for the mime type "application/vnd.ipunplugged.rcprofile".
	ApplicationVndIpunpluggedRcprofile = "application/vnd.ipunplugged.rcprofile"

	// ApplicationVndIrepositoryPackageXml is the constant for the mime type "application/vnd.irepository.package+xml".
	ApplicationVndIrepositoryPackageXml = "application/vnd.irepository.package+xml"

	// ApplicationVndIsXpr is the constant for the mime type "application/vnd.is-xpr".
	ApplicationVndIsXpr = "application/vnd.is-xpr"

	// ApplicationVndIsacFcs is the constant for the mime type "application/vnd.isac.fcs".
	ApplicationVndIsacFcs = "application/vnd.isac.fcs"

	// ApplicationVndJam is the constant for the mime type "application/vnd.jam".
	ApplicationVndJam = "application/vnd.jam"

	// ApplicationVndIso1178310Zip is the constant for the mime type "application/vnd.iso11783-10+zip".
	ApplicationVndIso1178310Zip = "application/vnd.iso11783-10+zip"

	// ApplicationVndJapannetDirectoryService is the constant for the mime type "application/vnd.japannet-directory-service".
	ApplicationVndJapannetDirectoryService = "application/vnd.japannet-directory-service"

	// ApplicationVndJapannetJpnstoreWakeup is the constant for the mime type "application/vnd.japannet-jpnstore-wakeup".
	ApplicationVndJapannetJpnstoreWakeup = "application/vnd.japannet-jpnstore-wakeup"

	// ApplicationVndJapannetPaymentWakeup is the constant for the mime type "application/vnd.japannet-payment-wakeup".
	ApplicationVndJapannetPaymentWakeup = "application/vnd.japannet-payment-wakeup"

	// ApplicationVndJapannetRegistration is the constant for the mime type "application/vnd.japannet-registration".
	ApplicationVndJapannetRegistration = "application/vnd.japannet-registration"

	// ApplicationVndJapannetRegistrationWakeup is the constant for the mime type "application/vnd.japannet-registration-wakeup".
	ApplicationVndJapannetRegistrationWakeup = "application/vnd.japannet-registration-wakeup"

	// ApplicationVndJapannetSetstoreWakeup is the constant for the mime type "application/vnd.japannet-setstore-wakeup".
	ApplicationVndJapannetSetstoreWakeup = "application/vnd.japannet-setstore-wakeup"

	// ApplicationVndJapannetVerification is the constant for the mime type "application/vnd.japannet-verification".
	ApplicationVndJapannetVerification = "application/vnd.japannet-verification"

	// ApplicationVndJapannetVerificationWakeup is the constant for the mime type "application/vnd.japannet-verification-wakeup".
	ApplicationVndJapannetVerificationWakeup = "application/vnd.japannet-verification-wakeup"

	// ApplicationVndJcpJavameMidletRms is the constant for the mime type "application/vnd.jcp.javame.midlet-rms".
	ApplicationVndJcpJavameMidletRms = "application/vnd.jcp.javame.midlet-rms"

	// ApplicationVndJisp is the constant for the mime type "application/vnd.jisp".
	ApplicationVndJisp = "application/vnd.jisp"

	// ApplicationVndJoostJodaArchive is the constant for the mime type "application/vnd.joost.joda-archive".
	ApplicationVndJoostJodaArchive = "application/vnd.joost.joda-archive"

	// ApplicationVndJskIsdnNgn is the constant for the mime type "application/vnd.jsk.isdn-ngn".
	ApplicationVndJskIsdnNgn = "application/vnd.jsk.isdn-ngn"

	// ApplicationVndKahootz is the constant for the mime type "application/vnd.kahootz".
	ApplicationVndKahootz = "application/vnd.kahootz"

	// ApplicationVndKdeKarbon is the constant for the mime type "application/vnd.kde.karbon".
	ApplicationVndKdeKarbon = "application/vnd.kde.karbon"

	// ApplicationVndKdeKchart is the constant for the mime type "application/vnd.kde.kchart".
	ApplicationVndKdeKchart = "application/vnd.kde.kchart"

	// ApplicationVndKdeKformula is the constant for the mime type "application/vnd.kde.kformula".
	ApplicationVndKdeKformula = "application/vnd.kde.kformula"

	// ApplicationVndKdeKivio is the constant for the mime type "application/vnd.kde.kivio".
	ApplicationVndKdeKivio = "application/vnd.kde.kivio"

	// ApplicationVndKdeKontour is the constant for the mime type "application/vnd.kde.kontour".
	ApplicationVndKdeKontour = "application/vnd.kde.kontour"

	// ApplicationVndKdeKpresenter is the constant for the mime type "application/vnd.kde.kpresenter".
	ApplicationVndKdeKpresenter = "application/vnd.kde.kpresenter"

	// ApplicationVndKdeKspread is the constant for the mime type "application/vnd.kde.kspread".
	ApplicationVndKdeKspread = "application/vnd.kde.kspread"

	// ApplicationVndKdeKword is the constant for the mime type "application/vnd.kde.kword".
	ApplicationVndKdeKword = "application/vnd.kde.kword"

	// ApplicationVndKenameaapp is the constant for the mime type "application/vnd.kenameaapp".
	ApplicationVndKenameaapp = "application/vnd.kenameaapp"

	// ApplicationVndKidspiration is the constant for the mime type "application/vnd.kidspiration".
	ApplicationVndKidspiration = "application/vnd.kidspiration"

	// ApplicationVndKinar is the constant for the mime type "application/vnd.Kinar".
	ApplicationVndKinar = "application/vnd.Kinar"

	// ApplicationVndKoan is the constant for the mime type "application/vnd.koan".
	ApplicationVndKoan = "application/vnd.koan"

	// ApplicationVndKodakDescriptor is the constant for the mime type "application/vnd.kodak-descriptor".
	ApplicationVndKodakDescriptor = "application/vnd.kodak-descriptor"

	// ApplicationVndLas is the constant for the mime type "application/vnd.las".
	ApplicationVndLas = "application/vnd.las"

	// ApplicationVndLasLasJson is the constant for the mime type "application/vnd.las.las+json".
	ApplicationVndLasLasJson = "application/vnd.las.las+json"

	// ApplicationVndLasLasXml is the constant for the mime type "application/vnd.las.las+xml".
	ApplicationVndLasLasXml = "application/vnd.las.las+xml"

	// ApplicationVndLaszip is the constant for the mime type "application/vnd.laszip".
	ApplicationVndLaszip = "application/vnd.laszip"

	// ApplicationVndLeapJson is the constant for the mime type "application/vnd.leap+json".
	ApplicationVndLeapJson = "application/vnd.leap+json"

	// ApplicationVndLibertyRequestXml is the constant for the mime type "application/vnd.liberty-request+xml".
	ApplicationVndLibertyRequestXml = "application/vnd.liberty-request+xml"

	// ApplicationVndLlamagraphicsLifeBalanceDesktop is the constant for the mime type "application/vnd.llamagraphics.life-balance.desktop".
	ApplicationVndLlamagraphicsLifeBalanceDesktop = "application/vnd.llamagraphics.life-balance.desktop"

	// ApplicationVndLlamagraphicsLifeBalanceExchangeXml is the constant for the mime type "application/vnd.llamagraphics.life-balance.exchange+xml".
	ApplicationVndLlamagraphicsLifeBalanceExchangeXml = "application/vnd.llamagraphics.life-balance.exchange+xml"

	// ApplicationVndLogipipeCircuitZip is the constant for the mime type "application/vnd.logipipe.circuit+zip".
	ApplicationVndLogipipeCircuitZip = "application/vnd.logipipe.circuit+zip"

	// ApplicationVndLoom is the constant for the mime type "application/vnd.loom".
	ApplicationVndLoom = "application/vnd.loom"

	// ApplicationVndLotus123 is the constant for the mime type "application/vnd.lotus-1-2-3".
	ApplicationVndLotus123 = "application/vnd.lotus-1-2-3"

	// ApplicationVndLotusApproach is the constant for the mime type "application/vnd.lotus-approach".
	ApplicationVndLotusApproach = "application/vnd.lotus-approach"

	// ApplicationVndLotusFreelance is the constant for the mime type "application/vnd.lotus-freelance".
	ApplicationVndLotusFreelance = "application/vnd.lotus-freelance"

	// ApplicationVndLotusNotes is the constant for the mime type "application/vnd.lotus-notes".
	ApplicationVndLotusNotes = "application/vnd.lotus-notes"

	// ApplicationVndLotusOrganizer is the constant for the mime type "application/vnd.lotus-organizer".
	ApplicationVndLotusOrganizer = "application/vnd.lotus-organizer"

	// ApplicationVndLotusScreencam is the constant for the mime type "application/vnd.lotus-screencam".
	ApplicationVndLotusScreencam = "application/vnd.lotus-screencam"

	// ApplicationVndLotusWordpro is the constant for the mime type "application/vnd.lotus-wordpro".
	ApplicationVndLotusWordpro = "application/vnd.lotus-wordpro"

	// ApplicationVndMacportsPortpkg is the constant for the mime type "application/vnd.macports.portpkg".
	ApplicationVndMacportsPortpkg = "application/vnd.macports.portpkg"

	// ApplicationVndMapboxVectorTile is the constant for the mime type "application/vnd.mapbox-vector-tile".
	ApplicationVndMapboxVectorTile = "application/vnd.mapbox-vector-tile"

	// ApplicationVndMarlinDrmActiontokenXml is the constant for the mime type "application/vnd.marlin.drm.actiontoken+xml".
	ApplicationVndMarlinDrmActiontokenXml = "application/vnd.marlin.drm.actiontoken+xml"

	// ApplicationVndMarlinDrmConftokenXml is the constant for the mime type "application/vnd.marlin.drm.conftoken+xml".
	ApplicationVndMarlinDrmConftokenXml = "application/vnd.marlin.drm.conftoken+xml"

	// ApplicationVndMarlinDrmLicenseXml is the constant for the mime type "application/vnd.marlin.drm.license+xml".
	ApplicationVndMarlinDrmLicenseXml = "application/vnd.marlin.drm.license+xml"

	// ApplicationVndMarlinDrmMdcf is the constant for the mime type "application/vnd.marlin.drm.mdcf".
	ApplicationVndMarlinDrmMdcf = "application/vnd.marlin.drm.mdcf"

	// ApplicationVndMasonJson is the constant for the mime type "application/vnd.mason+json".
	ApplicationVndMasonJson = "application/vnd.mason+json"

	// ApplicationVndMaxmindMaxmindDb is the constant for the mime type "application/vnd.maxmind.maxmind-db".
	ApplicationVndMaxmindMaxmindDb = "application/vnd.maxmind.maxmind-db"

	// ApplicationVndMcd is the constant for the mime type "application/vnd.mcd".
	ApplicationVndMcd = "application/vnd.mcd"

	// ApplicationVndMedcalcdata is the constant for the mime type "application/vnd.medcalcdata".
	ApplicationVndMedcalcdata = "application/vnd.medcalcdata"

	// ApplicationVndMediastationCdkey is the constant for the mime type "application/vnd.mediastation.cdkey".
	ApplicationVndMediastationCdkey = "application/vnd.mediastation.cdkey"

	// ApplicationVndMeridianSlingshot is the constant for the mime type "application/vnd.meridian-slingshot".
	ApplicationVndMeridianSlingshot = "application/vnd.meridian-slingshot"

	// ApplicationVndMFER is the constant for the mime type "application/vnd.MFER".
	ApplicationVndMFER = "application/vnd.MFER"

	// ApplicationVndMfmp is the constant for the mime type "application/vnd.mfmp".
	ApplicationVndMfmp = "application/vnd.mfmp"

	// ApplicationVndMicroJson is the constant for the mime type "application/vnd.micro+json".
	ApplicationVndMicroJson = "application/vnd.micro+json"

	// ApplicationVndMicrografxFlo is the constant for the mime type "application/vnd.micrografx.flo".
	ApplicationVndMicrografxFlo = "application/vnd.micrografx.flo"

	// ApplicationVndMicrografxIgx is the constant for the mime type "application/vnd.micrografx.igx".
	ApplicationVndMicrografxIgx = "application/vnd.micrografx.igx"

	// ApplicationVndMicrosoftPortableExecutable is the constant for the mime type "application/vnd.microsoft.portable-executable".
	ApplicationVndMicrosoftPortableExecutable = "application/vnd.microsoft.portable-executable"

	// ApplicationVndMicrosoftWindowsThumbnailCache is the constant for the mime type "application/vnd.microsoft.windows.thumbnail-cache".
	ApplicationVndMicrosoftWindowsThumbnailCache = "application/vnd.microsoft.windows.thumbnail-cache"

	// ApplicationVndMieleJson is the constant for the mime type "application/vnd.miele+json".
	ApplicationVndMieleJson = "application/vnd.miele+json"

	// ApplicationVndMif is the constant for the mime type "application/vnd.mif".
	ApplicationVndMif = "application/vnd.mif"

	// ApplicationVndMinisoftHp3000Save is the constant for the mime type "application/vnd.minisoft-hp3000-save".
	ApplicationVndMinisoftHp3000Save = "application/vnd.minisoft-hp3000-save"

	// ApplicationVndMitsubishiMistyGuardTrustweb is the constant for the mime type "application/vnd.mitsubishi.misty-guard.trustweb".
	ApplicationVndMitsubishiMistyGuardTrustweb = "application/vnd.mitsubishi.misty-guard.trustweb"

	// ApplicationVndMobiusDAF is the constant for the mime type "application/vnd.Mobius.DAF".
	ApplicationVndMobiusDAF = "application/vnd.Mobius.DAF"

	// ApplicationVndMobiusDIS is the constant for the mime type "application/vnd.Mobius.DIS".
	ApplicationVndMobiusDIS = "application/vnd.Mobius.DIS"

	// ApplicationVndMobiusMBK is the constant for the mime type "application/vnd.Mobius.MBK".
	ApplicationVndMobiusMBK = "application/vnd.Mobius.MBK"

	// ApplicationVndMobiusMQY is the constant for the mime type "application/vnd.Mobius.MQY".
	ApplicationVndMobiusMQY = "application/vnd.Mobius.MQY"

	// ApplicationVndMobiusMSL is the constant for the mime type "application/vnd.Mobius.MSL".
	ApplicationVndMobiusMSL = "application/vnd.Mobius.MSL"

	// ApplicationVndMobiusPLC is the constant for the mime type "application/vnd.Mobius.PLC".
	ApplicationVndMobiusPLC = "application/vnd.Mobius.PLC"

	// ApplicationVndMobiusTXF is the constant for the mime type "application/vnd.Mobius.TXF".
	ApplicationVndMobiusTXF = "application/vnd.Mobius.TXF"

	// ApplicationVndMophunApplication is the constant for the mime type "application/vnd.mophun.application".
	ApplicationVndMophunApplication = "application/vnd.mophun.application"

	// ApplicationVndMophunCertificate is the constant for the mime type "application/vnd.mophun.certificate".
	ApplicationVndMophunCertificate = "application/vnd.mophun.certificate"

	// ApplicationVndMotorolaFlexsuite is the constant for the mime type "application/vnd.motorola.flexsuite".
	ApplicationVndMotorolaFlexsuite = "application/vnd.motorola.flexsuite"

	// ApplicationVndMotorolaFlexsuiteAdsi is the constant for the mime type "application/vnd.motorola.flexsuite.adsi".
	ApplicationVndMotorolaFlexsuiteAdsi = "application/vnd.motorola.flexsuite.adsi"

	// ApplicationVndMotorolaFlexsuiteFis is the constant for the mime type "application/vnd.motorola.flexsuite.fis".
	ApplicationVndMotorolaFlexsuiteFis = "application/vnd.motorola.flexsuite.fis"

	// ApplicationVndMotorolaFlexsuiteGotap is the constant for the mime type "application/vnd.motorola.flexsuite.gotap".
	ApplicationVndMotorolaFlexsuiteGotap = "application/vnd.motorola.flexsuite.gotap"

	// ApplicationVndMotorolaFlexsuiteKmr is the constant for the mime type "application/vnd.motorola.flexsuite.kmr".
	ApplicationVndMotorolaFlexsuiteKmr = "application/vnd.motorola.flexsuite.kmr"

	// ApplicationVndMotorolaFlexsuiteTtc is the constant for the mime type "application/vnd.motorola.flexsuite.ttc".
	ApplicationVndMotorolaFlexsuiteTtc = "application/vnd.motorola.flexsuite.ttc"

	// ApplicationVndMotorolaFlexsuiteWem is the constant for the mime type "application/vnd.motorola.flexsuite.wem".
	ApplicationVndMotorolaFlexsuiteWem = "application/vnd.motorola.flexsuite.wem"

	// ApplicationVndMotorolaIprm is the constant for the mime type "application/vnd.motorola.iprm".
	ApplicationVndMotorolaIprm = "application/vnd.motorola.iprm"

	// ApplicationVndMozillaXulXml is the constant for the mime type "application/vnd.mozilla.xul+xml".
	ApplicationVndMozillaXulXml = "application/vnd.mozilla.xul+xml"

	// ApplicationVndMsArtgalry is the constant for the mime type "application/vnd.ms-artgalry".
	ApplicationVndMsArtgalry = "application/vnd.ms-artgalry"

	// ApplicationVndMsAsf is the constant for the mime type "application/vnd.ms-asf".
	ApplicationVndMsAsf = "application/vnd.ms-asf"

	// ApplicationVndMsCabCompressed is the constant for the mime type "application/vnd.ms-cab-compressed".
	ApplicationVndMsCabCompressed = "application/vnd.ms-cab-compressed"

	// ApplicationVndMs3mfdocument is the constant for the mime type "application/vnd.ms-3mfdocument".
	ApplicationVndMs3mfdocument = "application/vnd.ms-3mfdocument"

	// ApplicationVndMsExcel is the constant for the mime type "application/vnd.ms-excel".
	ApplicationVndMsExcel = "application/vnd.ms-excel"

	// ApplicationVndMsExcelAddinMacroEnabled12 is the constant for the mime type "application/vnd.ms-excel.addin.macroEnabled.12".
	ApplicationVndMsExcelAddinMacroEnabled12 = "application/vnd.ms-excel.addin.macroEnabled.12"

	// ApplicationVndMsExcelSheetBinaryMacroEnabled12 is the constant for the mime type "application/vnd.ms-excel.sheet.binary.macroEnabled.12".
	ApplicationVndMsExcelSheetBinaryMacroEnabled12 = "application/vnd.ms-excel.sheet.binary.macroEnabled.12"

	// ApplicationVndMsExcelSheetMacroEnabled12 is the constant for the mime type "application/vnd.ms-excel.sheet.macroEnabled.12".
	ApplicationVndMsExcelSheetMacroEnabled12 = "application/vnd.ms-excel.sheet.macroEnabled.12"

	// ApplicationVndMsExcelTemplateMacroEnabled12 is the constant for the mime type "application/vnd.ms-excel.template.macroEnabled.12".
	ApplicationVndMsExcelTemplateMacroEnabled12 = "application/vnd.ms-excel.template.macroEnabled.12"

	// ApplicationVndMsFontobject is the constant for the mime type "application/vnd.ms-fontobject".
	ApplicationVndMsFontobject = "application/vnd.ms-fontobject"

	// ApplicationVndMsHtmlhelp is the constant for the mime type "application/vnd.ms-htmlhelp".
	ApplicationVndMsHtmlhelp = "application/vnd.ms-htmlhelp"

	// ApplicationVndMsIms is the constant for the mime type "application/vnd.ms-ims".
	ApplicationVndMsIms = "application/vnd.ms-ims"

	// ApplicationVndMsLrm is the constant for the mime type "application/vnd.ms-lrm".
	ApplicationVndMsLrm = "application/vnd.ms-lrm"

	// ApplicationVndMsOfficeActiveXXml is the constant for the mime type "application/vnd.ms-office.activeX+xml".
	ApplicationVndMsOfficeActiveXXml = "application/vnd.ms-office.activeX+xml"

	// ApplicationVndMsOfficetheme is the constant for the mime type "application/vnd.ms-officetheme".
	ApplicationVndMsOfficetheme = "application/vnd.ms-officetheme"

	// ApplicationVndMsPlayreadyInitiatorXml is the constant for the mime type "application/vnd.ms-playready.initiator+xml".
	ApplicationVndMsPlayreadyInitiatorXml = "application/vnd.ms-playready.initiator+xml"

	// ApplicationVndMsPowerpoint is the constant for the mime type "application/vnd.ms-powerpoint".
	ApplicationVndMsPowerpoint = "application/vnd.ms-powerpoint"

	// ApplicationVndMsPowerpointAddinMacroEnabled12 is the constant for the mime type "application/vnd.ms-powerpoint.addin.macroEnabled.12".
	ApplicationVndMsPowerpointAddinMacroEnabled12 = "application/vnd.ms-powerpoint.addin.macroEnabled.12"

	// ApplicationVndMsPowerpointPresentationMacroEnabled12 is the constant for the mime type "application/vnd.ms-powerpoint.presentation.macroEnabled.12".
	ApplicationVndMsPowerpointPresentationMacroEnabled12 = "application/vnd.ms-powerpoint.presentation.macroEnabled.12"

	// ApplicationVndMsPowerpointSlideMacroEnabled12 is the constant for the mime type "application/vnd.ms-powerpoint.slide.macroEnabled.12".
	ApplicationVndMsPowerpointSlideMacroEnabled12 = "application/vnd.ms-powerpoint.slide.macroEnabled.12"

	// ApplicationVndMsPowerpointSlideshowMacroEnabled12 is the constant for the mime type "application/vnd.ms-powerpoint.slideshow.macroEnabled.12".
	ApplicationVndMsPowerpointSlideshowMacroEnabled12 = "application/vnd.ms-powerpoint.slideshow.macroEnabled.12"

	// ApplicationVndMsPowerpointTemplateMacroEnabled12 is the constant for the mime type "application/vnd.ms-powerpoint.template.macroEnabled.12".
	ApplicationVndMsPowerpointTemplateMacroEnabled12 = "application/vnd.ms-powerpoint.template.macroEnabled.12"

	// ApplicationVndMsPrintDeviceCapabilitiesXml is the constant for the mime type "application/vnd.ms-PrintDeviceCapabilities+xml".
	ApplicationVndMsPrintDeviceCapabilitiesXml = "application/vnd.ms-PrintDeviceCapabilities+xml"

	// ApplicationVndMsPrintSchemaTicketXml is the constant for the mime type "application/vnd.ms-PrintSchemaTicket+xml".
	ApplicationVndMsPrintSchemaTicketXml = "application/vnd.ms-PrintSchemaTicket+xml"

	// ApplicationVndMsProject is the constant for the mime type "application/vnd.ms-project".
	ApplicationVndMsProject = "application/vnd.ms-project"

	// ApplicationVndMsTnef is the constant for the mime type "application/vnd.ms-tnef".
	ApplicationVndMsTnef = "application/vnd.ms-tnef"

	// ApplicationVndMsWindowsDevicepairing is the constant for the mime type "application/vnd.ms-windows.devicepairing".
	ApplicationVndMsWindowsDevicepairing = "application/vnd.ms-windows.devicepairing"

	// ApplicationVndMsWindowsNwprintingOob is the constant for the mime type "application/vnd.ms-windows.nwprinting.oob".
	ApplicationVndMsWindowsNwprintingOob = "application/vnd.ms-windows.nwprinting.oob"

	// ApplicationVndMsWindowsPrinterpairing is the constant for the mime type "application/vnd.ms-windows.printerpairing".
	ApplicationVndMsWindowsPrinterpairing = "application/vnd.ms-windows.printerpairing"

	// ApplicationVndMsWindowsWsdOob is the constant for the mime type "application/vnd.ms-windows.wsd.oob".
	ApplicationVndMsWindowsWsdOob = "application/vnd.ms-windows.wsd.oob"

	// ApplicationVndMsWmdrmLicChlgReq is the constant for the mime type "application/vnd.ms-wmdrm.lic-chlg-req".
	ApplicationVndMsWmdrmLicChlgReq = "application/vnd.ms-wmdrm.lic-chlg-req"

	// ApplicationVndMsWmdrmLicResp is the constant for the mime type "application/vnd.ms-wmdrm.lic-resp".
	ApplicationVndMsWmdrmLicResp = "application/vnd.ms-wmdrm.lic-resp"

	// ApplicationVndMsWmdrmMeterChlgReq is the constant for the mime type "application/vnd.ms-wmdrm.meter-chlg-req".
	ApplicationVndMsWmdrmMeterChlgReq = "application/vnd.ms-wmdrm.meter-chlg-req"

	// ApplicationVndMsWmdrmMeterResp is the constant for the mime type "application/vnd.ms-wmdrm.meter-resp".
	ApplicationVndMsWmdrmMeterResp = "application/vnd.ms-wmdrm.meter-resp"

	// ApplicationVndMsWordDocumentMacroEnabled12 is the constant for the mime type "application/vnd.ms-word.document.macroEnabled.12".
	ApplicationVndMsWordDocumentMacroEnabled12 = "application/vnd.ms-word.document.macroEnabled.12"

	// ApplicationVndMsWordTemplateMacroEnabled12 is the constant for the mime type "application/vnd.ms-word.template.macroEnabled.12".
	ApplicationVndMsWordTemplateMacroEnabled12 = "application/vnd.ms-word.template.macroEnabled.12"

	// ApplicationVndMsWorks is the constant for the mime type "application/vnd.ms-works".
	ApplicationVndMsWorks = "application/vnd.ms-works"

	// ApplicationVndMsWpl is the constant for the mime type "application/vnd.ms-wpl".
	ApplicationVndMsWpl = "application/vnd.ms-wpl"

	// ApplicationVndMsXpsdocument is the constant for the mime type "application/vnd.ms-xpsdocument".
	ApplicationVndMsXpsdocument = "application/vnd.ms-xpsdocument"

	// ApplicationVndMsaDiskImage is the constant for the mime type "application/vnd.msa-disk-image".
	ApplicationVndMsaDiskImage = "application/vnd.msa-disk-image"

	// ApplicationVndMseq is the constant for the mime type "application/vnd.mseq".
	ApplicationVndMseq = "application/vnd.mseq"

	// ApplicationVndMsign is the constant for the mime type "application/vnd.msign".
	ApplicationVndMsign = "application/vnd.msign"

	// ApplicationVndMultiadCreator is the constant for the mime type "application/vnd.multiad.creator".
	ApplicationVndMultiadCreator = "application/vnd.multiad.creator"

	// ApplicationVndMultiadCreatorCif is the constant for the mime type "application/vnd.multiad.creator.cif".
	ApplicationVndMultiadCreatorCif = "application/vnd.multiad.creator.cif"

	// ApplicationVndMusician is the constant for the mime type "application/vnd.musician".
	ApplicationVndMusician = "application/vnd.musician"

	// ApplicationVndMusicNiff is the constant for the mime type "application/vnd.music-niff".
	ApplicationVndMusicNiff = "application/vnd.music-niff"

	// ApplicationVndMuveeStyle is the constant for the mime type "application/vnd.muvee.style".
	ApplicationVndMuveeStyle = "application/vnd.muvee.style"

	// ApplicationVndMynfc is the constant for the mime type "application/vnd.mynfc".
	ApplicationVndMynfc = "application/vnd.mynfc"

	// ApplicationVndNcdControl is the constant for the mime type "application/vnd.ncd.control".
	ApplicationVndNcdControl = "application/vnd.ncd.control"

	// ApplicationVndNcdReference is the constant for the mime type "application/vnd.ncd.reference".
	ApplicationVndNcdReference = "application/vnd.ncd.reference"

	// ApplicationVndNearstInvJson is the constant for the mime type "application/vnd.nearst.inv+json".
	ApplicationVndNearstInvJson = "application/vnd.nearst.inv+json"

	// ApplicationVndNervana is the constant for the mime type "application/vnd.nervana".
	ApplicationVndNervana = "application/vnd.nervana"

	// ApplicationVndNetfpx is the constant for the mime type "application/vnd.netfpx".
	ApplicationVndNetfpx = "application/vnd.netfpx"

	// ApplicationVndNeurolanguageNlu is the constant for the mime type "application/vnd.neurolanguage.nlu".
	ApplicationVndNeurolanguageNlu = "application/vnd.neurolanguage.nlu"

	// ApplicationVndNimn is the constant for the mime type "application/vnd.nimn".
	ApplicationVndNimn = "application/vnd.nimn"

	// ApplicationVndNintendoSnesRom is the constant for the mime type "application/vnd.nintendo.snes.rom".
	ApplicationVndNintendoSnesRom = "application/vnd.nintendo.snes.rom"

	// ApplicationVndNintendoNitroRom is the constant for the mime type "application/vnd.nintendo.nitro.rom".
	ApplicationVndNintendoNitroRom = "application/vnd.nintendo.nitro.rom"

	// ApplicationVndNitf is the constant for the mime type "application/vnd.nitf".
	ApplicationVndNitf = "application/vnd.nitf"

	// ApplicationVndNoblenetDirectory is the constant for the mime type "application/vnd.noblenet-directory".
	ApplicationVndNoblenetDirectory = "application/vnd.noblenet-directory"

	// ApplicationVndNoblenetSealer is the constant for the mime type "application/vnd.noblenet-sealer".
	ApplicationVndNoblenetSealer = "application/vnd.noblenet-sealer"

	// ApplicationVndNoblenetWeb is the constant for the mime type "application/vnd.noblenet-web".
	ApplicationVndNoblenetWeb = "application/vnd.noblenet-web"

	// ApplicationVndNokiaCatalogs is the constant for the mime type "application/vnd.nokia.catalogs".
	ApplicationVndNokiaCatalogs = "application/vnd.nokia.catalogs"

	// ApplicationVndNokiaConmlWbxml is the constant for the mime type "application/vnd.nokia.conml+wbxml".
	ApplicationVndNokiaConmlWbxml = "application/vnd.nokia.conml+wbxml"

	// ApplicationVndNokiaConmlXml is the constant for the mime type "application/vnd.nokia.conml+xml".
	ApplicationVndNokiaConmlXml = "application/vnd.nokia.conml+xml"

	// ApplicationVndNokiaIptvConfigXml is the constant for the mime type "application/vnd.nokia.iptv.config+xml".
	ApplicationVndNokiaIptvConfigXml = "application/vnd.nokia.iptv.config+xml"

	// ApplicationVndNokiaISDSRadioPresets is the constant for the mime type "application/vnd.nokia.iSDS-radio-presets".
	ApplicationVndNokiaISDSRadioPresets = "application/vnd.nokia.iSDS-radio-presets"

	// ApplicationVndNokiaLandmarkWbxml is the constant for the mime type "application/vnd.nokia.landmark+wbxml".
	ApplicationVndNokiaLandmarkWbxml = "application/vnd.nokia.landmark+wbxml"

	// ApplicationVndNokiaLandmarkXml is the constant for the mime type "application/vnd.nokia.landmark+xml".
	ApplicationVndNokiaLandmarkXml = "application/vnd.nokia.landmark+xml"

	// ApplicationVndNokiaLandmarkcollectionXml is the constant for the mime type "application/vnd.nokia.landmarkcollection+xml".
	ApplicationVndNokiaLandmarkcollectionXml = "application/vnd.nokia.landmarkcollection+xml"

	// ApplicationVndNokiaNcd is the constant for the mime type "application/vnd.nokia.ncd".
	ApplicationVndNokiaNcd = "application/vnd.nokia.ncd"

	// ApplicationVndNokiaNGageAcXml is the constant for the mime type "application/vnd.nokia.n-gage.ac+xml".
	ApplicationVndNokiaNGageAcXml = "application/vnd.nokia.n-gage.ac+xml"

	// ApplicationVndNokiaNGageData is the constant for the mime type "application/vnd.nokia.n-gage.data".
	ApplicationVndNokiaNGageData = "application/vnd.nokia.n-gage.data"

	// ApplicationVndNokiaNGageSymbianInstall is the constant for the mime type "application/vnd.nokia.n-gage.symbian.install - OBSOLETE; no replacement given".
	ApplicationVndNokiaNGageSymbianInstall = "application/vnd.nokia.n-gage.symbian.install"

	// ApplicationVndNokiaPcdWbxml is the constant for the mime type "application/vnd.nokia.pcd+wbxml".
	ApplicationVndNokiaPcdWbxml = "application/vnd.nokia.pcd+wbxml"

	// ApplicationVndNokiaPcdXml is the constant for the mime type "application/vnd.nokia.pcd+xml".
	ApplicationVndNokiaPcdXml = "application/vnd.nokia.pcd+xml"

	// ApplicationVndNokiaRadioPreset is the constant for the mime type "application/vnd.nokia.radio-preset".
	ApplicationVndNokiaRadioPreset = "application/vnd.nokia.radio-preset"

	// ApplicationVndNokiaRadioPresets is the constant for the mime type "application/vnd.nokia.radio-presets".
	ApplicationVndNokiaRadioPresets = "application/vnd.nokia.radio-presets"

	// ApplicationVndNovadigmEDM is the constant for the mime type "application/vnd.novadigm.EDM".
	ApplicationVndNovadigmEDM = "application/vnd.novadigm.EDM"

	// ApplicationVndNovadigmEDX is the constant for the mime type "application/vnd.novadigm.EDX".
	ApplicationVndNovadigmEDX = "application/vnd.novadigm.EDX"

	// ApplicationVndNovadigmEXT is the constant for the mime type "application/vnd.novadigm.EXT".
	ApplicationVndNovadigmEXT = "application/vnd.novadigm.EXT"

	// ApplicationVndNttLocalContentShare is the constant for the mime type "application/vnd.ntt-local.content-share".
	ApplicationVndNttLocalContentShare = "application/vnd.ntt-local.content-share"

	// ApplicationVndNttLocalFileTransfer is the constant for the mime type "application/vnd.ntt-local.file-transfer".
	ApplicationVndNttLocalFileTransfer = "application/vnd.ntt-local.file-transfer"

	// ApplicationVndNttLocalOgwRemoteAccess is the constant for the mime type "application/vnd.ntt-local.ogw_remote-access".
	ApplicationVndNttLocalOgwRemoteAccess = "application/vnd.ntt-local.ogw_remote-access"

	// ApplicationVndNttLocalSipTaRemote is the constant for the mime type "application/vnd.ntt-local.sip-ta_remote".
	ApplicationVndNttLocalSipTaRemote = "application/vnd.ntt-local.sip-ta_remote"

	// ApplicationVndNttLocalSipTaTcpStream is the constant for the mime type "application/vnd.ntt-local.sip-ta_tcp_stream".
	ApplicationVndNttLocalSipTaTcpStream = "application/vnd.ntt-local.sip-ta_tcp_stream"

	// ApplicationVndOasisOpendocumentChart is the constant for the mime type "application/vnd.oasis.opendocument.chart".
	ApplicationVndOasisOpendocumentChart = "application/vnd.oasis.opendocument.chart"

	// ApplicationVndOasisOpendocumentChartTemplate is the constant for the mime type "application/vnd.oasis.opendocument.chart-template".
	ApplicationVndOasisOpendocumentChartTemplate = "application/vnd.oasis.opendocument.chart-template"

	// ApplicationVndOasisOpendocumentDatabase is the constant for the mime type "application/vnd.oasis.opendocument.database".
	ApplicationVndOasisOpendocumentDatabase = "application/vnd.oasis.opendocument.database"

	// ApplicationVndOasisOpendocumentFormula is the constant for the mime type "application/vnd.oasis.opendocument.formula".
	ApplicationVndOasisOpendocumentFormula = "application/vnd.oasis.opendocument.formula"

	// ApplicationVndOasisOpendocumentFormulaTemplate is the constant for the mime type "application/vnd.oasis.opendocument.formula-template".
	ApplicationVndOasisOpendocumentFormulaTemplate = "application/vnd.oasis.opendocument.formula-template"

	// ApplicationVndOasisOpendocumentGraphics is the constant for the mime type "application/vnd.oasis.opendocument.graphics".
	ApplicationVndOasisOpendocumentGraphics = "application/vnd.oasis.opendocument.graphics"

	// ApplicationVndOasisOpendocumentGraphicsTemplate is the constant for the mime type "application/vnd.oasis.opendocument.graphics-template".
	ApplicationVndOasisOpendocumentGraphicsTemplate = "application/vnd.oasis.opendocument.graphics-template"

	// ApplicationVndOasisOpendocumentImage is the constant for the mime type "application/vnd.oasis.opendocument.image".
	ApplicationVndOasisOpendocumentImage = "application/vnd.oasis.opendocument.image"

	// ApplicationVndOasisOpendocumentImageTemplate is the constant for the mime type "application/vnd.oasis.opendocument.image-template".
	ApplicationVndOasisOpendocumentImageTemplate = "application/vnd.oasis.opendocument.image-template"

	// ApplicationVndOasisOpendocumentPresentation is the constant for the mime type "application/vnd.oasis.opendocument.presentation".
	ApplicationVndOasisOpendocumentPresentation = "application/vnd.oasis.opendocument.presentation"

	// ApplicationVndOasisOpendocumentPresentationTemplate is the constant for the mime type "application/vnd.oasis.opendocument.presentation-template".
	ApplicationVndOasisOpendocumentPresentationTemplate = "application/vnd.oasis.opendocument.presentation-template"

	// ApplicationVndOasisOpendocumentSpreadsheet is the constant for the mime type "application/vnd.oasis.opendocument.spreadsheet".
	ApplicationVndOasisOpendocumentSpreadsheet = "application/vnd.oasis.opendocument.spreadsheet"

	// ApplicationVndOasisOpendocumentSpreadsheetTemplate is the constant for the mime type "application/vnd.oasis.opendocument.spreadsheet-template".
	ApplicationVndOasisOpendocumentSpreadsheetTemplate = "application/vnd.oasis.opendocument.spreadsheet-template"

	// ApplicationVndOasisOpendocumentText is the constant for the mime type "application/vnd.oasis.opendocument.text".
	ApplicationVndOasisOpendocumentText = "application/vnd.oasis.opendocument.text"

	// ApplicationVndOasisOpendocumentTextMaster is the constant for the mime type "application/vnd.oasis.opendocument.text-master".
	ApplicationVndOasisOpendocumentTextMaster = "application/vnd.oasis.opendocument.text-master"

	// ApplicationVndOasisOpendocumentTextTemplate is the constant for the mime type "application/vnd.oasis.opendocument.text-template".
	ApplicationVndOasisOpendocumentTextTemplate = "application/vnd.oasis.opendocument.text-template"

	// ApplicationVndOasisOpendocumentTextWeb is the constant for the mime type "application/vnd.oasis.opendocument.text-web".
	ApplicationVndOasisOpendocumentTextWeb = "application/vnd.oasis.opendocument.text-web"

	// ApplicationVndObn is the constant for the mime type "application/vnd.obn".
	ApplicationVndObn = "application/vnd.obn"

	// ApplicationVndOcfCbor is the constant for the mime type "application/vnd.ocf+cbor".
	ApplicationVndOcfCbor = "application/vnd.ocf+cbor"

	// ApplicationVndOftnL10NJson is the constant for the mime type "application/vnd.oftn.l10n+json".
	ApplicationVndOftnL10NJson = "application/vnd.oftn.l10n+json"

	// ApplicationVndOipfContentaccessdownloadXml is the constant for the mime type "application/vnd.oipf.contentaccessdownload+xml".
	ApplicationVndOipfContentaccessdownloadXml = "application/vnd.oipf.contentaccessdownload+xml"

	// ApplicationVndOipfContentaccessstreamingXml is the constant for the mime type "application/vnd.oipf.contentaccessstreaming+xml".
	ApplicationVndOipfContentaccessstreamingXml = "application/vnd.oipf.contentaccessstreaming+xml"

	// ApplicationVndOipfCspgHexbinary is the constant for the mime type "application/vnd.oipf.cspg-hexbinary".
	ApplicationVndOipfCspgHexbinary = "application/vnd.oipf.cspg-hexbinary"

	// ApplicationVndOipfDaeSvgXml is the constant for the mime type "application/vnd.oipf.dae.svg+xml".
	ApplicationVndOipfDaeSvgXml = "application/vnd.oipf.dae.svg+xml"

	// ApplicationVndOipfDaeXhtmlXml is the constant for the mime type "application/vnd.oipf.dae.xhtml+xml".
	ApplicationVndOipfDaeXhtmlXml = "application/vnd.oipf.dae.xhtml+xml"

	// ApplicationVndOipfMippvcontrolmessageXml is the constant for the mime type "application/vnd.oipf.mippvcontrolmessage+xml".
	ApplicationVndOipfMippvcontrolmessageXml = "application/vnd.oipf.mippvcontrolmessage+xml"

	// ApplicationVndOipfPaeGem is the constant for the mime type "application/vnd.oipf.pae.gem".
	ApplicationVndOipfPaeGem = "application/vnd.oipf.pae.gem"

	// ApplicationVndOipfSpdiscoveryXml is the constant for the mime type "application/vnd.oipf.spdiscovery+xml".
	ApplicationVndOipfSpdiscoveryXml = "application/vnd.oipf.spdiscovery+xml"

	// ApplicationVndOipfSpdlistXml is the constant for the mime type "application/vnd.oipf.spdlist+xml".
	ApplicationVndOipfSpdlistXml = "application/vnd.oipf.spdlist+xml"

	// ApplicationVndOipfUeprofileXml is the constant for the mime type "application/vnd.oipf.ueprofile+xml".
	ApplicationVndOipfUeprofileXml = "application/vnd.oipf.ueprofile+xml"

	// ApplicationVndOipfUserprofileXml is the constant for the mime type "application/vnd.oipf.userprofile+xml".
	ApplicationVndOipfUserprofileXml = "application/vnd.oipf.userprofile+xml"

	// ApplicationVndOlpcSugar is the constant for the mime type "application/vnd.olpc-sugar".
	ApplicationVndOlpcSugar = "application/vnd.olpc-sugar"

	// ApplicationVndOmaBcastAssociatedProcedureParameterXml is the constant for the mime type "application/vnd.oma.bcast.associated-procedure-parameter+xml".
	ApplicationVndOmaBcastAssociatedProcedureParameterXml = "application/vnd.oma.bcast.associated-procedure-parameter+xml"

	// ApplicationVndOmaBcastDrmTriggerXml is the constant for the mime type "application/vnd.oma.bcast.drm-trigger+xml".
	ApplicationVndOmaBcastDrmTriggerXml = "application/vnd.oma.bcast.drm-trigger+xml"

	// ApplicationVndOmaBcastImdXml is the constant for the mime type "application/vnd.oma.bcast.imd+xml".
	ApplicationVndOmaBcastImdXml = "application/vnd.oma.bcast.imd+xml"

	// ApplicationVndOmaBcastLtkm is the constant for the mime type "application/vnd.oma.bcast.ltkm".
	ApplicationVndOmaBcastLtkm = "application/vnd.oma.bcast.ltkm"

	// ApplicationVndOmaBcastNotificationXml is the constant for the mime type "application/vnd.oma.bcast.notification+xml".
	ApplicationVndOmaBcastNotificationXml = "application/vnd.oma.bcast.notification+xml"

	// ApplicationVndOmaBcastProvisioningtrigger is the constant for the mime type "application/vnd.oma.bcast.provisioningtrigger".
	ApplicationVndOmaBcastProvisioningtrigger = "application/vnd.oma.bcast.provisioningtrigger"

	// ApplicationVndOmaBcastSgboot is the constant for the mime type "application/vnd.oma.bcast.sgboot".
	ApplicationVndOmaBcastSgboot = "application/vnd.oma.bcast.sgboot"

	// ApplicationVndOmaBcastSgddXml is the constant for the mime type "application/vnd.oma.bcast.sgdd+xml".
	ApplicationVndOmaBcastSgddXml = "application/vnd.oma.bcast.sgdd+xml"

	// ApplicationVndOmaBcastSgdu is the constant for the mime type "application/vnd.oma.bcast.sgdu".
	ApplicationVndOmaBcastSgdu = "application/vnd.oma.bcast.sgdu"

	// ApplicationVndOmaBcastSimpleSymbolContainer is the constant for the mime type "application/vnd.oma.bcast.simple-symbol-container".
	ApplicationVndOmaBcastSimpleSymbolContainer = "application/vnd.oma.bcast.simple-symbol-container"

	// ApplicationVndOmaBcastSmartcardTriggerXml is the constant for the mime type "application/vnd.oma.bcast.smartcard-trigger+xml".
	ApplicationVndOmaBcastSmartcardTriggerXml = "application/vnd.oma.bcast.smartcard-trigger+xml"

	// ApplicationVndOmaBcastSprovXml is the constant for the mime type "application/vnd.oma.bcast.sprov+xml".
	ApplicationVndOmaBcastSprovXml = "application/vnd.oma.bcast.sprov+xml"

	// ApplicationVndOmaBcastStkm is the constant for the mime type "application/vnd.oma.bcast.stkm".
	ApplicationVndOmaBcastStkm = "application/vnd.oma.bcast.stkm"

	// ApplicationVndOmaCabAddressBookXml is the constant for the mime type "application/vnd.oma.cab-address-book+xml".
	ApplicationVndOmaCabAddressBookXml = "application/vnd.oma.cab-address-book+xml"

	// ApplicationVndOmaCabFeatureHandlerXml is the constant for the mime type "application/vnd.oma.cab-feature-handler+xml".
	ApplicationVndOmaCabFeatureHandlerXml = "application/vnd.oma.cab-feature-handler+xml"

	// ApplicationVndOmaCabPccXml is the constant for the mime type "application/vnd.oma.cab-pcc+xml".
	ApplicationVndOmaCabPccXml = "application/vnd.oma.cab-pcc+xml"

	// ApplicationVndOmaCabSubsInviteXml is the constant for the mime type "application/vnd.oma.cab-subs-invite+xml".
	ApplicationVndOmaCabSubsInviteXml = "application/vnd.oma.cab-subs-invite+xml"

	// ApplicationVndOmaCabUserPrefsXml is the constant for the mime type "application/vnd.oma.cab-user-prefs+xml".
	ApplicationVndOmaCabUserPrefsXml = "application/vnd.oma.cab-user-prefs+xml"

	// ApplicationVndOmaDcd is the constant for the mime type "application/vnd.oma.dcd".
	ApplicationVndOmaDcd = "application/vnd.oma.dcd"

	// ApplicationVndOmaDcdc is the constant for the mime type "application/vnd.oma.dcdc".
	ApplicationVndOmaDcdc = "application/vnd.oma.dcdc"

	// ApplicationVndOmaDd2Xml is the constant for the mime type "application/vnd.oma.dd2+xml".
	ApplicationVndOmaDd2Xml = "application/vnd.oma.dd2+xml"

	// ApplicationVndOmaDrmRisdXml is the constant for the mime type "application/vnd.oma.drm.risd+xml".
	ApplicationVndOmaDrmRisdXml = "application/vnd.oma.drm.risd+xml"

	// ApplicationVndOmaGroupUsageListXml is the constant for the mime type "application/vnd.oma.group-usage-list+xml".
	ApplicationVndOmaGroupUsageListXml = "application/vnd.oma.group-usage-list+xml"

	// ApplicationVndOmaLwm2MJson is the constant for the mime type "application/vnd.oma.lwm2m+json".
	ApplicationVndOmaLwm2MJson = "application/vnd.oma.lwm2m+json"

	// ApplicationVndOmaLwm2MTlv is the constant for the mime type "application/vnd.oma.lwm2m+tlv".
	ApplicationVndOmaLwm2MTlv = "application/vnd.oma.lwm2m+tlv"

	// ApplicationVndOmaPalXml is the constant for the mime type "application/vnd.oma.pal+xml".
	ApplicationVndOmaPalXml = "application/vnd.oma.pal+xml"

	// ApplicationVndOmaPocDetailedProgressReportXml is the constant for the mime type "application/vnd.oma.poc.detailed-progress-report+xml".
	ApplicationVndOmaPocDetailedProgressReportXml = "application/vnd.oma.poc.detailed-progress-report+xml"

	// ApplicationVndOmaPocFinalReportXml is the constant for the mime type "application/vnd.oma.poc.final-report+xml".
	ApplicationVndOmaPocFinalReportXml = "application/vnd.oma.poc.final-report+xml"

	// ApplicationVndOmaPocGroupsXml is the constant for the mime type "application/vnd.oma.poc.groups+xml".
	ApplicationVndOmaPocGroupsXml = "application/vnd.oma.poc.groups+xml"

	// ApplicationVndOmaPocInvocationDescriptorXml is the constant for the mime type "application/vnd.oma.poc.invocation-descriptor+xml".
	ApplicationVndOmaPocInvocationDescriptorXml = "application/vnd.oma.poc.invocation-descriptor+xml"

	// ApplicationVndOmaPocOptimizedProgressReportXml is the constant for the mime type "application/vnd.oma.poc.optimized-progress-report+xml".
	ApplicationVndOmaPocOptimizedProgressReportXml = "application/vnd.oma.poc.optimized-progress-report+xml"

	// ApplicationVndOmaPush is the constant for the mime type "application/vnd.oma.push".
	ApplicationVndOmaPush = "application/vnd.oma.push"

	// ApplicationVndOmaScidmMessagesXml is the constant for the mime type "application/vnd.oma.scidm.messages+xml".
	ApplicationVndOmaScidmMessagesXml = "application/vnd.oma.scidm.messages+xml"

	// ApplicationVndOmaXcapDirectoryXml is the constant for the mime type "application/vnd.oma.xcap-directory+xml".
	ApplicationVndOmaXcapDirectoryXml = "application/vnd.oma.xcap-directory+xml"

	// ApplicationVndOmadsEmailXml is the constant for the mime type "application/vnd.omads-email+xml".
	ApplicationVndOmadsEmailXml = "application/vnd.omads-email+xml"

	// ApplicationVndOmadsFileXml is the constant for the mime type "application/vnd.omads-file+xml".
	ApplicationVndOmadsFileXml = "application/vnd.omads-file+xml"

	// ApplicationVndOmadsFolderXml is the constant for the mime type "application/vnd.omads-folder+xml".
	ApplicationVndOmadsFolderXml = "application/vnd.omads-folder+xml"

	// ApplicationVndOmalocSuplInit is the constant for the mime type "application/vnd.omaloc-supl-init".
	ApplicationVndOmalocSuplInit = "application/vnd.omaloc-supl-init"

	// ApplicationVndOmaScwsConfig is the constant for the mime type "application/vnd.oma-scws-config".
	ApplicationVndOmaScwsConfig = "application/vnd.oma-scws-config"

	// ApplicationVndOmaScwsHttpRequest is the constant for the mime type "application/vnd.oma-scws-http-request".
	ApplicationVndOmaScwsHttpRequest = "application/vnd.oma-scws-http-request"

	// ApplicationVndOmaScwsHttpResponse is the constant for the mime type "application/vnd.oma-scws-http-response".
	ApplicationVndOmaScwsHttpResponse = "application/vnd.oma-scws-http-response"

	// ApplicationVndOnepager is the constant for the mime type "application/vnd.onepager".
	ApplicationVndOnepager = "application/vnd.onepager"

	// ApplicationVndOnepagertamp is the constant for the mime type "application/vnd.onepagertamp".
	ApplicationVndOnepagertamp = "application/vnd.onepagertamp"

	// ApplicationVndOnepagertamx is the constant for the mime type "application/vnd.onepagertamx".
	ApplicationVndOnepagertamx = "application/vnd.onepagertamx"

	// ApplicationVndOnepagertat is the constant for the mime type "application/vnd.onepagertat".
	ApplicationVndOnepagertat = "application/vnd.onepagertat"

	// ApplicationVndOnepagertatp is the constant for the mime type "application/vnd.onepagertatp".
	ApplicationVndOnepagertatp = "application/vnd.onepagertatp"

	// ApplicationVndOnepagertatx is the constant for the mime type "application/vnd.onepagertatx".
	ApplicationVndOnepagertatx = "application/vnd.onepagertatx"

	// ApplicationVndOpenbloxGameBinary is the constant for the mime type "application/vnd.openblox.game-binary".
	ApplicationVndOpenbloxGameBinary = "application/vnd.openblox.game-binary"

	// ApplicationVndOpenbloxGameXml is the constant for the mime type "application/vnd.openblox.game+xml".
	ApplicationVndOpenbloxGameXml = "application/vnd.openblox.game+xml"

	// ApplicationVndOpeneyeOeb is the constant for the mime type "application/vnd.openeye.oeb".
	ApplicationVndOpeneyeOeb = "application/vnd.openeye.oeb"

	// ApplicationVndOpenstreetmapDataXml is the constant for the mime type "application/vnd.openstreetmap.data+xml".
	ApplicationVndOpenstreetmapDataXml = "application/vnd.openstreetmap.data+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentCustomPropertiesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.custom-properties+xml".
	ApplicationVndOpenxmlformatsOfficedocumentCustomPropertiesXml = "application/vnd.openxmlformats-officedocument.custom-properties+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentCustomXmlPropertiesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.customXmlProperties+xml".
	ApplicationVndOpenxmlformatsOfficedocumentCustomXmlPropertiesXml = "application/vnd.openxmlformats-officedocument.customXmlProperties+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawing+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingXml = "application/vnd.openxmlformats-officedocument.drawing+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingmlChartXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawingml.chart+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingmlChartXml = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingmlChartshapesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawingml.chartshapes+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingmlChartshapesXml = "application/vnd.openxmlformats-officedocument.drawingml.chartshapes+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramColorsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawingml.diagramColors+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramColorsXml = "application/vnd.openxmlformats-officedocument.drawingml.diagramColors+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramDataXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawingml.diagramData+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramDataXml = "application/vnd.openxmlformats-officedocument.drawingml.diagramData+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramLayoutXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramLayoutXml = "application/vnd.openxmlformats-officedocument.drawingml.diagramLayout+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramStyleXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.drawingml.diagramStyle+xml".
	ApplicationVndOpenxmlformatsOfficedocumentDrawingmlDiagramStyleXml = "application/vnd.openxmlformats-officedocument.drawingml.diagramStyle+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentExtendedPropertiesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.extended-properties+xml".
	ApplicationVndOpenxmlformatsOfficedocumentExtendedPropertiesXml = "application/vnd.openxmlformats-officedocument.extended-properties+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlCommentAuthorsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.commentAuthors+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlCommentAuthorsXml = "application/vnd.openxmlformats-officedocument.presentationml.commentAuthors+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlCommentsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.comments+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlCommentsXml = "application/vnd.openxmlformats-officedocument.presentationml.comments+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlHandoutMasterXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.handoutMaster+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlHandoutMasterXml = "application/vnd.openxmlformats-officedocument.presentationml.handoutMaster+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlNotesMasterXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlNotesMasterXml = "application/vnd.openxmlformats-officedocument.presentationml.notesMaster+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlNotesSlideXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlNotesSlideXml = "application/vnd.openxmlformats-officedocument.presentationml.notesSlide+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.presentation".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentation = "application/vnd.openxmlformats-officedocument.presentationml.presentation"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentationMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresentationMainXml = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresPropsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.presProps+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlPresPropsXml = "application/vnd.openxmlformats-officedocument.presentationml.presProps+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlide is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slide".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlide = "application/vnd.openxmlformats-officedocument.presentationml.slide"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slide+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideXml = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideLayoutXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideLayoutXml = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideMasterXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideMasterXml = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideshow is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slideshow".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideshow = "application/vnd.openxmlformats-officedocument.presentationml.slideshow"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideshowMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideshowMainXml = "application/vnd.openxmlformats-officedocument.presentationml.slideshow.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideUpdateInfoXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.slideUpdateInfo+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlSlideUpdateInfoXml = "application/vnd.openxmlformats-officedocument.presentationml.slideUpdateInfo+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTableStylesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTableStylesXml = "application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTagsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.tags+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTagsXml = "application/vnd.openxmlformats-officedocument.presentationml.tags+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTemplate is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.template".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTemplate = "application/vnd.openxmlformats-officedocument.presentationml.template"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTemplateMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.template.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlTemplateMainXml = "application/vnd.openxmlformats-officedocument.presentationml.template.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentPresentationmlViewPropsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml".
	ApplicationVndOpenxmlformatsOfficedocumentPresentationmlViewPropsXml = "application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlCalcChainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.calcChain+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlCalcChainXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.calcChain+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlChartsheetXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlChartsheetXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlCommentsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlCommentsXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlConnectionsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.connections+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlConnectionsXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.connections+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlDialogsheetXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.dialogsheet+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlDialogsheetXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.dialogsheet+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlExternalLinkXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.externalLink+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlExternalLinkXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.externalLink+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlPivotCacheDefinitionXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheDefinition+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlPivotCacheDefinitionXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheDefinition+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlPivotCacheRecordsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheRecords+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlPivotCacheRecordsXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheRecords+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlPivotTableXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotTable+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlPivotTableXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotTable+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlQueryTableXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.queryTable+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlQueryTableXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.queryTable+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlRevisionHeadersXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.revisionHeaders+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlRevisionHeadersXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.revisionHeaders+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlRevisionLogXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.revisionLog+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlRevisionLogXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.revisionLog+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSharedStringsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSharedStringsXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheet = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheetMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheetMainXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheetMetadataXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.sheetMetadata+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlSheetMetadataXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheetMetadata+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlStylesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlStylesXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTableXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTableXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTableSingleCellsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.tableSingleCells+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTableSingleCellsXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.tableSingleCells+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTemplate is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.template".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTemplate = "application/vnd.openxmlformats-officedocument.spreadsheetml.template"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTemplateMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlTemplateMainXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlUserNamesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.userNames+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlUserNamesXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.userNames+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlVolatileDependenciesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.volatileDependencies+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlVolatileDependenciesXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.volatileDependencies+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlWorksheetXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml".
	ApplicationVndOpenxmlformatsOfficedocumentSpreadsheetmlWorksheetXml = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentThemeXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.theme+xml".
	ApplicationVndOpenxmlformatsOfficedocumentThemeXml = "application/vnd.openxmlformats-officedocument.theme+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentThemeOverrideXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.themeOverride+xml".
	ApplicationVndOpenxmlformatsOfficedocumentThemeOverrideXml = "application/vnd.openxmlformats-officedocument.themeOverride+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentVmlDrawing is the constant for the mime type "application/vnd.openxmlformats-officedocument.vmlDrawing".
	ApplicationVndOpenxmlformatsOfficedocumentVmlDrawing = "application/vnd.openxmlformats-officedocument.vmlDrawing"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlCommentsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlCommentsXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.document".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocumentGlossaryXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.document.glossary+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocumentGlossaryXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.glossary+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocumentMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocumentMainXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlEndnotesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlEndnotesXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.endnotes+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlFontTableXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlFontTableXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlFooterXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlFooterXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlFootnotesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlFootnotesXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.footnotes+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlNumberingXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlNumberingXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlSettingsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlSettingsXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlStylesXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlStylesXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlTemplate is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.template".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlTemplate = "application/vnd.openxmlformats-officedocument.wordprocessingml.template"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlTemplateMainXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlTemplateMainXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.template.main+xml"

	// ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlWebSettingsXml is the constant for the mime type "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml".
	ApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlWebSettingsXml = "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"

	// ApplicationVndOpenxmlformatsPackageCorePropertiesXml is the constant for the mime type "application/vnd.openxmlformats-package.core-properties+xml".
	ApplicationVndOpenxmlformatsPackageCorePropertiesXml = "application/vnd.openxmlformats-package.core-properties+xml"

	// ApplicationVndOpenxmlformatsPackageDigitalSignatureXmlsignatureXml is the constant for the mime type "application/vnd.openxmlformats-package.digital-signature-xmlsignature+xml".
	ApplicationVndOpenxmlformatsPackageDigitalSignatureXmlsignatureXml = "application/vnd.openxmlformats-package.digital-signature-xmlsignature+xml"

	// ApplicationVndOpenxmlformatsPackageRelationshipsXml is the constant for the mime type "application/vnd.openxmlformats-package.relationships+xml".
	ApplicationVndOpenxmlformatsPackageRelationshipsXml = "application/vnd.openxmlformats-package.relationships+xml"

	// ApplicationVndOracleResourceJson is the constant for the mime type "application/vnd.oracle.resource+json".
	ApplicationVndOracleResourceJson = "application/vnd.oracle.resource+json"

	// ApplicationVndOrangeIndata is the constant for the mime type "application/vnd.orange.indata".
	ApplicationVndOrangeIndata = "application/vnd.orange.indata"

	// ApplicationVndOsaNetdeploy is the constant for the mime type "application/vnd.osa.netdeploy".
	ApplicationVndOsaNetdeploy = "application/vnd.osa.netdeploy"

	// ApplicationVndOsgeoMapguidePackage is the constant for the mime type "application/vnd.osgeo.mapguide.package".
	ApplicationVndOsgeoMapguidePackage = "application/vnd.osgeo.mapguide.package"

	// ApplicationVndOsgiBundle is the constant for the mime type "application/vnd.osgi.bundle".
	ApplicationVndOsgiBundle = "application/vnd.osgi.bundle"

	// ApplicationVndOsgiDp is the constant for the mime type "application/vnd.osgi.dp".
	ApplicationVndOsgiDp = "application/vnd.osgi.dp"

	// ApplicationVndOsgiSubsystem is the constant for the mime type "application/vnd.osgi.subsystem".
	ApplicationVndOsgiSubsystem = "application/vnd.osgi.subsystem"

	// ApplicationVndOtpsCtKipXml is the constant for the mime type "application/vnd.otps.ct-kip+xml".
	ApplicationVndOtpsCtKipXml = "application/vnd.otps.ct-kip+xml"

	// ApplicationVndOxliCountgraph is the constant for the mime type "application/vnd.oxli.countgraph".
	ApplicationVndOxliCountgraph = "application/vnd.oxli.countgraph"

	// ApplicationVndPagerdutyJson is the constant for the mime type "application/vnd.pagerduty+json".
	ApplicationVndPagerdutyJson = "application/vnd.pagerduty+json"

	// ApplicationVndPalm is the constant for the mime type "application/vnd.palm".
	ApplicationVndPalm = "application/vnd.palm"

	// ApplicationVndPanoply is the constant for the mime type "application/vnd.panoply".
	ApplicationVndPanoply = "application/vnd.panoply"

	// ApplicationVndPaosXml is the constant for the mime type "application/vnd.paos.xml".
	ApplicationVndPaosXml = "application/vnd.paos.xml"

	// ApplicationVndPatentdive is the constant for the mime type "application/vnd.patentdive".
	ApplicationVndPatentdive = "application/vnd.patentdive"

	// ApplicationVndPatientecommsdoc is the constant for the mime type "application/vnd.patientecommsdoc".
	ApplicationVndPatientecommsdoc = "application/vnd.patientecommsdoc"

	// ApplicationVndPawaafile is the constant for the mime type "application/vnd.pawaafile".
	ApplicationVndPawaafile = "application/vnd.pawaafile"

	// ApplicationVndPcos is the constant for the mime type "application/vnd.pcos".
	ApplicationVndPcos = "application/vnd.pcos"

	// ApplicationVndPgFormat is the constant for the mime type "application/vnd.pg.format".
	ApplicationVndPgFormat = "application/vnd.pg.format"

	// ApplicationVndPgOsasli is the constant for the mime type "application/vnd.pg.osasli".
	ApplicationVndPgOsasli = "application/vnd.pg.osasli"

	// ApplicationVndPiaccessApplicationLicence is the constant for the mime type "application/vnd.piaccess.application-licence".
	ApplicationVndPiaccessApplicationLicence = "application/vnd.piaccess.application-licence"

	// ApplicationVndPicsel is the constant for the mime type "application/vnd.picsel".
	ApplicationVndPicsel = "application/vnd.picsel"

	// ApplicationVndPmiWidget is the constant for the mime type "application/vnd.pmi.widget".
	ApplicationVndPmiWidget = "application/vnd.pmi.widget"

	// ApplicationVndPocGroupAdvertisementXml is the constant for the mime type "application/vnd.poc.group-advertisement+xml".
	ApplicationVndPocGroupAdvertisementXml = "application/vnd.poc.group-advertisement+xml"

	// ApplicationVndPocketlearn is the constant for the mime type "application/vnd.pocketlearn".
	ApplicationVndPocketlearn = "application/vnd.pocketlearn"

	// ApplicationVndPowerbuilder6 is the constant for the mime type "application/vnd.powerbuilder6".
	ApplicationVndPowerbuilder6 = "application/vnd.powerbuilder6"

	// ApplicationVndPowerbuilder6S is the constant for the mime type "application/vnd.powerbuilder6-s".
	ApplicationVndPowerbuilder6S = "application/vnd.powerbuilder6-s"

	// ApplicationVndPowerbuilder7 is the constant for the mime type "application/vnd.powerbuilder7".
	ApplicationVndPowerbuilder7 = "application/vnd.powerbuilder7"

	// ApplicationVndPowerbuilder75 is the constant for the mime type "application/vnd.powerbuilder75".
	ApplicationVndPowerbuilder75 = "application/vnd.powerbuilder75"

	// ApplicationVndPowerbuilder75S is the constant for the mime type "application/vnd.powerbuilder75-s".
	ApplicationVndPowerbuilder75S = "application/vnd.powerbuilder75-s"

	// ApplicationVndPowerbuilder7S is the constant for the mime type "application/vnd.powerbuilder7-s".
	ApplicationVndPowerbuilder7S = "application/vnd.powerbuilder7-s"

	// ApplicationVndPreminet is the constant for the mime type "application/vnd.preminet".
	ApplicationVndPreminet = "application/vnd.preminet"

	// ApplicationVndPreviewsystemsBox is the constant for the mime type "application/vnd.previewsystems.box".
	ApplicationVndPreviewsystemsBox = "application/vnd.previewsystems.box"

	// ApplicationVndProteusMagazine is the constant for the mime type "application/vnd.proteus.magazine".
	ApplicationVndProteusMagazine = "application/vnd.proteus.magazine"

	// ApplicationVndPsfs is the constant for the mime type "application/vnd.psfs".
	ApplicationVndPsfs = "application/vnd.psfs"

	// ApplicationVndPublishareDeltaTree is the constant for the mime type "application/vnd.publishare-delta-tree".
	ApplicationVndPublishareDeltaTree = "application/vnd.publishare-delta-tree"

	// ApplicationVndPviPtid1 is the constant for the mime type "application/vnd.pvi.ptid1".
	ApplicationVndPviPtid1 = "application/vnd.pvi.ptid1"

	// ApplicationVndPwgMultiplexed is the constant for the mime type "application/vnd.pwg-multiplexed".
	ApplicationVndPwgMultiplexed = "application/vnd.pwg-multiplexed"

	// ApplicationVndPwgXhtmlPrintXml is the constant for the mime type "application/vnd.pwg-xhtml-print+xml".
	ApplicationVndPwgXhtmlPrintXml = "application/vnd.pwg-xhtml-print+xml"

	// ApplicationVndQualcommBrewAppRes is the constant for the mime type "application/vnd.qualcomm.brew-app-res".
	ApplicationVndQualcommBrewAppRes = "application/vnd.qualcomm.brew-app-res"

	// ApplicationVndQuarantainenet is the constant for the mime type "application/vnd.quarantainenet".
	ApplicationVndQuarantainenet = "application/vnd.quarantainenet"

	// ApplicationVndQuarkQuarkXPress is the constant for the mime type "application/vnd.Quark.QuarkXPress".
	ApplicationVndQuarkQuarkXPress = "application/vnd.Quark.QuarkXPress"

	// ApplicationVndQuobjectQuoxdocument is the constant for the mime type "application/vnd.quobject-quoxdocument".
	ApplicationVndQuobjectQuoxdocument = "application/vnd.quobject-quoxdocument"

	// ApplicationVndRadisysMomlXml is the constant for the mime type "application/vnd.radisys.moml+xml".
	ApplicationVndRadisysMomlXml = "application/vnd.radisys.moml+xml"

	// ApplicationVndRadisysMsmlAuditConfXml is the constant for the mime type "application/vnd.radisys.msml-audit-conf+xml".
	ApplicationVndRadisysMsmlAuditConfXml = "application/vnd.radisys.msml-audit-conf+xml"

	// ApplicationVndRadisysMsmlAuditConnXml is the constant for the mime type "application/vnd.radisys.msml-audit-conn+xml".
	ApplicationVndRadisysMsmlAuditConnXml = "application/vnd.radisys.msml-audit-conn+xml"

	// ApplicationVndRadisysMsmlAuditDialogXml is the constant for the mime type "application/vnd.radisys.msml-audit-dialog+xml".
	ApplicationVndRadisysMsmlAuditDialogXml = "application/vnd.radisys.msml-audit-dialog+xml"

	// ApplicationVndRadisysMsmlAuditStreamXml is the constant for the mime type "application/vnd.radisys.msml-audit-stream+xml".
	ApplicationVndRadisysMsmlAuditStreamXml = "application/vnd.radisys.msml-audit-stream+xml"

	// ApplicationVndRadisysMsmlAuditXml is the constant for the mime type "application/vnd.radisys.msml-audit+xml".
	ApplicationVndRadisysMsmlAuditXml = "application/vnd.radisys.msml-audit+xml"

	// ApplicationVndRadisysMsmlConfXml is the constant for the mime type "application/vnd.radisys.msml-conf+xml".
	ApplicationVndRadisysMsmlConfXml = "application/vnd.radisys.msml-conf+xml"

	// ApplicationVndRadisysMsmlDialogBaseXml is the constant for the mime type "application/vnd.radisys.msml-dialog-base+xml".
	ApplicationVndRadisysMsmlDialogBaseXml = "application/vnd.radisys.msml-dialog-base+xml"

	// ApplicationVndRadisysMsmlDialogFaxDetectXml is the constant for the mime type "application/vnd.radisys.msml-dialog-fax-detect+xml".
	ApplicationVndRadisysMsmlDialogFaxDetectXml = "application/vnd.radisys.msml-dialog-fax-detect+xml"

	// ApplicationVndRadisysMsmlDialogFaxSendrecvXml is the constant for the mime type "application/vnd.radisys.msml-dialog-fax-sendrecv+xml".
	ApplicationVndRadisysMsmlDialogFaxSendrecvXml = "application/vnd.radisys.msml-dialog-fax-sendrecv+xml"

	// ApplicationVndRadisysMsmlDialogGroupXml is the constant for the mime type "application/vnd.radisys.msml-dialog-group+xml".
	ApplicationVndRadisysMsmlDialogGroupXml = "application/vnd.radisys.msml-dialog-group+xml"

	// ApplicationVndRadisysMsmlDialogSpeechXml is the constant for the mime type "application/vnd.radisys.msml-dialog-speech+xml".
	ApplicationVndRadisysMsmlDialogSpeechXml = "application/vnd.radisys.msml-dialog-speech+xml"

	// ApplicationVndRadisysMsmlDialogTransformXml is the constant for the mime type "application/vnd.radisys.msml-dialog-transform+xml".
	ApplicationVndRadisysMsmlDialogTransformXml = "application/vnd.radisys.msml-dialog-transform+xml"

	// ApplicationVndRadisysMsmlDialogXml is the constant for the mime type "application/vnd.radisys.msml-dialog+xml".
	ApplicationVndRadisysMsmlDialogXml = "application/vnd.radisys.msml-dialog+xml"

	// ApplicationVndRadisysMsmlXml is the constant for the mime type "application/vnd.radisys.msml+xml".
	ApplicationVndRadisysMsmlXml = "application/vnd.radisys.msml+xml"

	// ApplicationVndRainstorData is the constant for the mime type "application/vnd.rainstor.data".
	ApplicationVndRainstorData = "application/vnd.rainstor.data"

	// ApplicationVndRapid is the constant for the mime type "application/vnd.rapid".
	ApplicationVndRapid = "application/vnd.rapid"

	// ApplicationVndRar is the constant for the mime type "application/vnd.rar".
	ApplicationVndRar = "application/vnd.rar"

	// ApplicationVndRealvncBed is the constant for the mime type "application/vnd.realvnc.bed".
	ApplicationVndRealvncBed = "application/vnd.realvnc.bed"

	// ApplicationVndRecordareMusicxml is the constant for the mime type "application/vnd.recordare.musicxml".
	ApplicationVndRecordareMusicxml = "application/vnd.recordare.musicxml"

	// ApplicationVndRecordareMusicxmlXml is the constant for the mime type "application/vnd.recordare.musicxml+xml".
	ApplicationVndRecordareMusicxmlXml = "application/vnd.recordare.musicxml+xml"

	// ApplicationVndRenLearnRlprint is the constant for the mime type "application/vnd.RenLearn.rlprint".
	ApplicationVndRenLearnRlprint = "application/vnd.RenLearn.rlprint"

	// ApplicationVndRestfulJson is the constant for the mime type "application/vnd.restful+json".
	ApplicationVndRestfulJson = "application/vnd.restful+json"

	// ApplicationVndRigCryptonote is the constant for the mime type "application/vnd.rig.cryptonote".
	ApplicationVndRigCryptonote = "application/vnd.rig.cryptonote"

	// ApplicationVndRoute66Link66Xml is the constant for the mime type "application/vnd.route66.link66+xml".
	ApplicationVndRoute66Link66Xml = "application/vnd.route66.link66+xml"

	// ApplicationVndRs274x is the constant for the mime type "application/vnd.rs-274x".
	ApplicationVndRs274x = "application/vnd.rs-274x"

	// ApplicationVndRuckusDownload is the constant for the mime type "application/vnd.ruckus.download".
	ApplicationVndRuckusDownload = "application/vnd.ruckus.download"

	// ApplicationVndS3Sms is the constant for the mime type "application/vnd.s3sms".
	ApplicationVndS3Sms = "application/vnd.s3sms"

	// ApplicationVndSailingtrackerTrack is the constant for the mime type "application/vnd.sailingtracker.track".
	ApplicationVndSailingtrackerTrack = "application/vnd.sailingtracker.track"

	// ApplicationVndSbmCid is the constant for the mime type "application/vnd.sbm.cid".
	ApplicationVndSbmCid = "application/vnd.sbm.cid"

	// ApplicationVndSbmMid2 is the constant for the mime type "application/vnd.sbm.mid2".
	ApplicationVndSbmMid2 = "application/vnd.sbm.mid2"

	// ApplicationVndScribus is the constant for the mime type "application/vnd.scribus".
	ApplicationVndScribus = "application/vnd.scribus"

	// ApplicationVndSealed3df is the constant for the mime type "application/vnd.sealed.3df".
	ApplicationVndSealed3df = "application/vnd.sealed.3df"

	// ApplicationVndSealedCsf is the constant for the mime type "application/vnd.sealed.csf".
	ApplicationVndSealedCsf = "application/vnd.sealed.csf"

	// ApplicationVndSealedDoc is the constant for the mime type "application/vnd.sealed.doc".
	ApplicationVndSealedDoc = "application/vnd.sealed.doc"

	// ApplicationVndSealedEml is the constant for the mime type "application/vnd.sealed.eml".
	ApplicationVndSealedEml = "application/vnd.sealed.eml"

	// ApplicationVndSealedMht is the constant for the mime type "application/vnd.sealed.mht".
	ApplicationVndSealedMht = "application/vnd.sealed.mht"

	// ApplicationVndSealedNet is the constant for the mime type "application/vnd.sealed.net".
	ApplicationVndSealedNet = "application/vnd.sealed.net"

	// ApplicationVndSealedPpt is the constant for the mime type "application/vnd.sealed.ppt".
	ApplicationVndSealedPpt = "application/vnd.sealed.ppt"

	// ApplicationVndSealedTiff is the constant for the mime type "application/vnd.sealed.tiff".
	ApplicationVndSealedTiff = "application/vnd.sealed.tiff"

	// ApplicationVndSealedXls is the constant for the mime type "application/vnd.sealed.xls".
	ApplicationVndSealedXls = "application/vnd.sealed.xls"

	// ApplicationVndSealedmediaSoftsealHtml is the constant for the mime type "application/vnd.sealedmedia.softseal.html".
	ApplicationVndSealedmediaSoftsealHtml = "application/vnd.sealedmedia.softseal.html"

	// ApplicationVndSealedmediaSoftsealPdf is the constant for the mime type "application/vnd.sealedmedia.softseal.pdf".
	ApplicationVndSealedmediaSoftsealPdf = "application/vnd.sealedmedia.softseal.pdf"

	// ApplicationVndSeemail is the constant for the mime type "application/vnd.seemail".
	ApplicationVndSeemail = "application/vnd.seemail"

	// ApplicationVndSema is the constant for the mime type "application/vnd.sema".
	ApplicationVndSema = "application/vnd.sema"

	// ApplicationVndSemd is the constant for the mime type "application/vnd.semd".
	ApplicationVndSemd = "application/vnd.semd"

	// ApplicationVndSemf is the constant for the mime type "application/vnd.semf".
	ApplicationVndSemf = "application/vnd.semf"

	// ApplicationVndShadeSaveFile is the constant for the mime type "application/vnd.shade-save-file".
	ApplicationVndShadeSaveFile = "application/vnd.shade-save-file"

	// ApplicationVndShanaInformedFormdata is the constant for the mime type "application/vnd.shana.informed.formdata".
	ApplicationVndShanaInformedFormdata = "application/vnd.shana.informed.formdata"

	// ApplicationVndShanaInformedFormtemplate is the constant for the mime type "application/vnd.shana.informed.formtemplate".
	ApplicationVndShanaInformedFormtemplate = "application/vnd.shana.informed.formtemplate"

	// ApplicationVndShanaInformedInterchange is the constant for the mime type "application/vnd.shana.informed.interchange".
	ApplicationVndShanaInformedInterchange = "application/vnd.shana.informed.interchange"

	// ApplicationVndShanaInformedPackage is the constant for the mime type "application/vnd.shana.informed.package".
	ApplicationVndShanaInformedPackage = "application/vnd.shana.informed.package"

	// ApplicationVndShootproofJson is the constant for the mime type "application/vnd.shootproof+json".
	ApplicationVndShootproofJson = "application/vnd.shootproof+json"

	// ApplicationVndShopkickJson is the constant for the mime type "application/vnd.shopkick+json".
	ApplicationVndShopkickJson = "application/vnd.shopkick+json"

	// ApplicationVndSigrokSession is the constant for the mime type "application/vnd.sigrok.session".
	ApplicationVndSigrokSession = "application/vnd.sigrok.session"

	// ApplicationVndSimTechMindMapper is the constant for the mime type "application/vnd.SimTech-MindMapper".
	ApplicationVndSimTechMindMapper = "application/vnd.SimTech-MindMapper"

	// ApplicationVndSirenJson is the constant for the mime type "application/vnd.siren+json".
	ApplicationVndSirenJson = "application/vnd.siren+json"

	// ApplicationVndSmaf is the constant for the mime type "application/vnd.smaf".
	ApplicationVndSmaf = "application/vnd.smaf"

	// ApplicationVndSmartNotebook is the constant for the mime type "application/vnd.smart.notebook".
	ApplicationVndSmartNotebook = "application/vnd.smart.notebook"

	// ApplicationVndSmartTeacher is the constant for the mime type "application/vnd.smart.teacher".
	ApplicationVndSmartTeacher = "application/vnd.smart.teacher"

	// ApplicationVndSoftware602FillerFormXml is the constant for the mime type "application/vnd.software602.filler.form+xml".
	ApplicationVndSoftware602FillerFormXml = "application/vnd.software602.filler.form+xml"

	// ApplicationVndSoftware602FillerFormXmlZip is the constant for the mime type "application/vnd.software602.filler.form-xml-zip".
	ApplicationVndSoftware602FillerFormXmlZip = "application/vnd.software602.filler.form-xml-zip"

	// ApplicationVndSolentSdkmXml is the constant for the mime type "application/vnd.solent.sdkm+xml".
	ApplicationVndSolentSdkmXml = "application/vnd.solent.sdkm+xml"

	// ApplicationVndSpotfireDxp is the constant for the mime type "application/vnd.spotfire.dxp".
	ApplicationVndSpotfireDxp = "application/vnd.spotfire.dxp"

	// ApplicationVndSpotfireSfs is the constant for the mime type "application/vnd.spotfire.sfs".
	ApplicationVndSpotfireSfs = "application/vnd.spotfire.sfs"

	// ApplicationVndSqlite3 is the constant for the mime type "application/vnd.sqlite3".
	ApplicationVndSqlite3 = "application/vnd.sqlite3"

	// ApplicationVndSssCod is the constant for the mime type "application/vnd.sss-cod".
	ApplicationVndSssCod = "application/vnd.sss-cod"

	// ApplicationVndSssDtf is the constant for the mime type "application/vnd.sss-dtf".
	ApplicationVndSssDtf = "application/vnd.sss-dtf"

	// ApplicationVndSssNtf is the constant for the mime type "application/vnd.sss-ntf".
	ApplicationVndSssNtf = "application/vnd.sss-ntf"

	// ApplicationVndStepmaniaPackage is the constant for the mime type "application/vnd.stepmania.package".
	ApplicationVndStepmaniaPackage = "application/vnd.stepmania.package"

	// ApplicationVndStepmaniaStepchart is the constant for the mime type "application/vnd.stepmania.stepchart".
	ApplicationVndStepmaniaStepchart = "application/vnd.stepmania.stepchart"

	// ApplicationVndStreetStream is the constant for the mime type "application/vnd.street-stream".
	ApplicationVndStreetStream = "application/vnd.street-stream"

	// ApplicationVndSunWadlXml is the constant for the mime type "application/vnd.sun.wadl+xml".
	ApplicationVndSunWadlXml = "application/vnd.sun.wadl+xml"

	// ApplicationVndSusCalendar is the constant for the mime type "application/vnd.sus-calendar".
	ApplicationVndSusCalendar = "application/vnd.sus-calendar"

	// ApplicationVndSvd is the constant for the mime type "application/vnd.svd".
	ApplicationVndSvd = "application/vnd.svd"

	// ApplicationVndSwiftviewIcs is the constant for the mime type "application/vnd.swiftview-ics".
	ApplicationVndSwiftviewIcs = "application/vnd.swiftview-ics"

	// ApplicationVndSyncmlDmNotification is the constant for the mime type "application/vnd.syncml.dm.notification".
	ApplicationVndSyncmlDmNotification = "application/vnd.syncml.dm.notification"

	// ApplicationVndSyncmlDmddfXml is the constant for the mime type "application/vnd.syncml.dmddf+xml".
	ApplicationVndSyncmlDmddfXml = "application/vnd.syncml.dmddf+xml"

	// ApplicationVndSyncmlDmtndsWbxml is the constant for the mime type "application/vnd.syncml.dmtnds+wbxml".
	ApplicationVndSyncmlDmtndsWbxml = "application/vnd.syncml.dmtnds+wbxml"

	// ApplicationVndSyncmlDmtndsXml is the constant for the mime type "application/vnd.syncml.dmtnds+xml".
	ApplicationVndSyncmlDmtndsXml = "application/vnd.syncml.dmtnds+xml"

	// ApplicationVndSyncmlDmddfWbxml is the constant for the mime type "application/vnd.syncml.dmddf+wbxml".
	ApplicationVndSyncmlDmddfWbxml = "application/vnd.syncml.dmddf+wbxml"

	// ApplicationVndSyncmlDmWbxml is the constant for the mime type "application/vnd.syncml.dm+wbxml".
	ApplicationVndSyncmlDmWbxml = "application/vnd.syncml.dm+wbxml"

	// ApplicationVndSyncmlDmXml is the constant for the mime type "application/vnd.syncml.dm+xml".
	ApplicationVndSyncmlDmXml = "application/vnd.syncml.dm+xml"

	// ApplicationVndSyncmlDsNotification is the constant for the mime type "application/vnd.syncml.ds.notification".
	ApplicationVndSyncmlDsNotification = "application/vnd.syncml.ds.notification"

	// ApplicationVndSyncmlXml is the constant for the mime type "application/vnd.syncml+xml".
	ApplicationVndSyncmlXml = "application/vnd.syncml+xml"

	// ApplicationVndTableschemaJson is the constant for the mime type "application/vnd.tableschema+json".
	ApplicationVndTableschemaJson = "application/vnd.tableschema+json"

	// ApplicationVndTaoIntentModuleArchive is the constant for the mime type "application/vnd.tao.intent-module-archive".
	ApplicationVndTaoIntentModuleArchive = "application/vnd.tao.intent-module-archive"

	// ApplicationVndTcpdumpPcap is the constant for the mime type "application/vnd.tcpdump.pcap".
	ApplicationVndTcpdumpPcap = "application/vnd.tcpdump.pcap"

	// ApplicationVndThinkCellPpttcJson is the constant for the mime type "application/vnd.think-cell.ppttc+json".
	ApplicationVndThinkCellPpttcJson = "application/vnd.think-cell.ppttc+json"

	// ApplicationVndTml is the constant for the mime type "application/vnd.tml".
	ApplicationVndTml = "application/vnd.tml"

	// ApplicationVndTmdMediaflexApiXml is the constant for the mime type "application/vnd.tmd.mediaflex.api+xml".
	ApplicationVndTmdMediaflexApiXml = "application/vnd.tmd.mediaflex.api+xml"

	// ApplicationVndTmobileLivetv is the constant for the mime type "application/vnd.tmobile-livetv".
	ApplicationVndTmobileLivetv = "application/vnd.tmobile-livetv"

	// ApplicationVndTriOnesource is the constant for the mime type "application/vnd.tri.onesource".
	ApplicationVndTriOnesource = "application/vnd.tri.onesource"

	// ApplicationVndTridTpt is the constant for the mime type "application/vnd.trid.tpt".
	ApplicationVndTridTpt = "application/vnd.trid.tpt"

	// ApplicationVndTriscapeMxs is the constant for the mime type "application/vnd.triscape.mxs".
	ApplicationVndTriscapeMxs = "application/vnd.triscape.mxs"

	// ApplicationVndTrueapp is the constant for the mime type "application/vnd.trueapp".
	ApplicationVndTrueapp = "application/vnd.trueapp"

	// ApplicationVndTruedoc is the constant for the mime type "application/vnd.truedoc".
	ApplicationVndTruedoc = "application/vnd.truedoc"

	// ApplicationVndUbisoftWebplayer is the constant for the mime type "application/vnd.ubisoft.webplayer".
	ApplicationVndUbisoftWebplayer = "application/vnd.ubisoft.webplayer"

	// ApplicationVndUfdl is the constant for the mime type "application/vnd.ufdl".
	ApplicationVndUfdl = "application/vnd.ufdl"

	// ApplicationVndUiqTheme is the constant for the mime type "application/vnd.uiq.theme".
	ApplicationVndUiqTheme = "application/vnd.uiq.theme"

	// ApplicationVndUmajin is the constant for the mime type "application/vnd.umajin".
	ApplicationVndUmajin = "application/vnd.umajin"

	// ApplicationVndUnity is the constant for the mime type "application/vnd.unity".
	ApplicationVndUnity = "application/vnd.unity"

	// ApplicationVndUomlXml is the constant for the mime type "application/vnd.uoml+xml".
	ApplicationVndUomlXml = "application/vnd.uoml+xml"

	// ApplicationVndUplanetAlert is the constant for the mime type "application/vnd.uplanet.alert".
	ApplicationVndUplanetAlert = "application/vnd.uplanet.alert"

	// ApplicationVndUplanetAlertWbxml is the constant for the mime type "application/vnd.uplanet.alert-wbxml".
	ApplicationVndUplanetAlertWbxml = "application/vnd.uplanet.alert-wbxml"

	// ApplicationVndUplanetBearerChoice is the constant for the mime type "application/vnd.uplanet.bearer-choice".
	ApplicationVndUplanetBearerChoice = "application/vnd.uplanet.bearer-choice"

	// ApplicationVndUplanetBearerChoiceWbxml is the constant for the mime type "application/vnd.uplanet.bearer-choice-wbxml".
	ApplicationVndUplanetBearerChoiceWbxml = "application/vnd.uplanet.bearer-choice-wbxml"

	// ApplicationVndUplanetCacheop is the constant for the mime type "application/vnd.uplanet.cacheop".
	ApplicationVndUplanetCacheop = "application/vnd.uplanet.cacheop"

	// ApplicationVndUplanetCacheopWbxml is the constant for the mime type "application/vnd.uplanet.cacheop-wbxml".
	ApplicationVndUplanetCacheopWbxml = "application/vnd.uplanet.cacheop-wbxml"

	// ApplicationVndUplanetChannel is the constant for the mime type "application/vnd.uplanet.channel".
	ApplicationVndUplanetChannel = "application/vnd.uplanet.channel"

	// ApplicationVndUplanetChannelWbxml is the constant for the mime type "application/vnd.uplanet.channel-wbxml".
	ApplicationVndUplanetChannelWbxml = "application/vnd.uplanet.channel-wbxml"

	// ApplicationVndUplanetList is the constant for the mime type "application/vnd.uplanet.list".
	ApplicationVndUplanetList = "application/vnd.uplanet.list"

	// ApplicationVndUplanetListcmd is the constant for the mime type "application/vnd.uplanet.listcmd".
	ApplicationVndUplanetListcmd = "application/vnd.uplanet.listcmd"

	// ApplicationVndUplanetListcmdWbxml is the constant for the mime type "application/vnd.uplanet.listcmd-wbxml".
	ApplicationVndUplanetListcmdWbxml = "application/vnd.uplanet.listcmd-wbxml"

	// ApplicationVndUplanetListWbxml is the constant for the mime type "application/vnd.uplanet.list-wbxml".
	ApplicationVndUplanetListWbxml = "application/vnd.uplanet.list-wbxml"

	// ApplicationVndUriMap is the constant for the mime type "application/vnd.uri-map".
	ApplicationVndUriMap = "application/vnd.uri-map"

	// ApplicationVndUplanetSignal is the constant for the mime type "application/vnd.uplanet.signal".
	ApplicationVndUplanetSignal = "application/vnd.uplanet.signal"

	// ApplicationVndValveSourceMaterial is the constant for the mime type "application/vnd.valve.source.material".
	ApplicationVndValveSourceMaterial = "application/vnd.valve.source.material"

	// ApplicationVndVcx is the constant for the mime type "application/vnd.vcx".
	ApplicationVndVcx = "application/vnd.vcx"

	// ApplicationVndVdStudy is the constant for the mime type "application/vnd.vd-study".
	ApplicationVndVdStudy = "application/vnd.vd-study"

	// ApplicationVndVectorworks is the constant for the mime type "application/vnd.vectorworks".
	ApplicationVndVectorworks = "application/vnd.vectorworks"

	// ApplicationVndVelJson is the constant for the mime type "application/vnd.vel+json".
	ApplicationVndVelJson = "application/vnd.vel+json"

	// ApplicationVndVerimatrixVcas is the constant for the mime type "application/vnd.verimatrix.vcas".
	ApplicationVndVerimatrixVcas = "application/vnd.verimatrix.vcas"

	// ApplicationVndVeryantThin is the constant for the mime type "application/vnd.veryant.thin".
	ApplicationVndVeryantThin = "application/vnd.veryant.thin"

	// ApplicationVndVesEncrypted is the constant for the mime type "application/vnd.ves.encrypted".
	ApplicationVndVesEncrypted = "application/vnd.ves.encrypted"

	// ApplicationVndVidsoftVidconference is the constant for the mime type "application/vnd.vidsoft.vidconference".
	ApplicationVndVidsoftVidconference = "application/vnd.vidsoft.vidconference"

	// ApplicationVndVisio is the constant for the mime type "application/vnd.visio".
	ApplicationVndVisio = "application/vnd.visio"

	// ApplicationVndVisionary is the constant for the mime type "application/vnd.visionary".
	ApplicationVndVisionary = "application/vnd.visionary"

	// ApplicationVndVividenceScriptfile is the constant for the mime type "application/vnd.vividence.scriptfile".
	ApplicationVndVividenceScriptfile = "application/vnd.vividence.scriptfile"

	// ApplicationVndVsf is the constant for the mime type "application/vnd.vsf".
	ApplicationVndVsf = "application/vnd.vsf"

	// ApplicationVndWapSic is the constant for the mime type "application/vnd.wap.sic".
	ApplicationVndWapSic = "application/vnd.wap.sic"

	// ApplicationVndWapSlc is the constant for the mime type "application/vnd.wap.slc".
	ApplicationVndWapSlc = "application/vnd.wap.slc"

	// ApplicationVndWapWbxml is the constant for the mime type "application/vnd.wap.wbxml".
	ApplicationVndWapWbxml = "application/vnd.wap.wbxml"

	// ApplicationVndWapWmlc is the constant for the mime type "application/vnd.wap.wmlc".
	ApplicationVndWapWmlc = "application/vnd.wap.wmlc"

	// ApplicationVndWapWmlscriptc is the constant for the mime type "application/vnd.wap.wmlscriptc".
	ApplicationVndWapWmlscriptc = "application/vnd.wap.wmlscriptc"

	// ApplicationVndWebturbo is the constant for the mime type "application/vnd.webturbo".
	ApplicationVndWebturbo = "application/vnd.webturbo"

	// ApplicationVndWfaP2P is the constant for the mime type "application/vnd.wfa.p2p".
	ApplicationVndWfaP2P = "application/vnd.wfa.p2p"

	// ApplicationVndWfaWsc is the constant for the mime type "application/vnd.wfa.wsc".
	ApplicationVndWfaWsc = "application/vnd.wfa.wsc"

	// ApplicationVndWindowsDevicepairing is the constant for the mime type "application/vnd.windows.devicepairing".
	ApplicationVndWindowsDevicepairing = "application/vnd.windows.devicepairing"

	// ApplicationVndWmc is the constant for the mime type "application/vnd.wmc".
	ApplicationVndWmc = "application/vnd.wmc"

	// ApplicationVndWmfBootstrap is the constant for the mime type "application/vnd.wmf.bootstrap".
	ApplicationVndWmfBootstrap = "application/vnd.wmf.bootstrap"

	// ApplicationVndWolframMathematica is the constant for the mime type "application/vnd.wolfram.mathematica".
	ApplicationVndWolframMathematica = "application/vnd.wolfram.mathematica"

	// ApplicationVndWolframMathematicaPackage is the constant for the mime type "application/vnd.wolfram.mathematica.package".
	ApplicationVndWolframMathematicaPackage = "application/vnd.wolfram.mathematica.package"

	// ApplicationVndWolframPlayer is the constant for the mime type "application/vnd.wolfram.player".
	ApplicationVndWolframPlayer = "application/vnd.wolfram.player"

	// ApplicationVndWordperfect is the constant for the mime type "application/vnd.wordperfect".
	ApplicationVndWordperfect = "application/vnd.wordperfect"

	// ApplicationVndWqd is the constant for the mime type "application/vnd.wqd".
	ApplicationVndWqd = "application/vnd.wqd"

	// ApplicationVndWrqHp3000Labelled is the constant for the mime type "application/vnd.wrq-hp3000-labelled".
	ApplicationVndWrqHp3000Labelled = "application/vnd.wrq-hp3000-labelled"

	// ApplicationVndWtStf is the constant for the mime type "application/vnd.wt.stf".
	ApplicationVndWtStf = "application/vnd.wt.stf"

	// ApplicationVndWvCspXml is the constant for the mime type "application/vnd.wv.csp+xml".
	ApplicationVndWvCspXml = "application/vnd.wv.csp+xml"

	// ApplicationVndWvCspWbxml is the constant for the mime type "application/vnd.wv.csp+wbxml".
	ApplicationVndWvCspWbxml = "application/vnd.wv.csp+wbxml"

	// ApplicationVndWvSspXml is the constant for the mime type "application/vnd.wv.ssp+xml".
	ApplicationVndWvSspXml = "application/vnd.wv.ssp+xml"

	// ApplicationVndXacmlJson is the constant for the mime type "application/vnd.xacml+json".
	ApplicationVndXacmlJson = "application/vnd.xacml+json"

	// ApplicationVndXara is the constant for the mime type "application/vnd.xara".
	ApplicationVndXara = "application/vnd.xara"

	// ApplicationVndXfdl is the constant for the mime type "application/vnd.xfdl".
	ApplicationVndXfdl = "application/vnd.xfdl"

	// ApplicationVndXfdlWebform is the constant for the mime type "application/vnd.xfdl.webform".
	ApplicationVndXfdlWebform = "application/vnd.xfdl.webform"

	// ApplicationVndXmiXml is the constant for the mime type "application/vnd.xmi+xml".
	ApplicationVndXmiXml = "application/vnd.xmi+xml"

	// ApplicationVndXmpieCpkg is the constant for the mime type "application/vnd.xmpie.cpkg".
	ApplicationVndXmpieCpkg = "application/vnd.xmpie.cpkg"

	// ApplicationVndXmpieDpkg is the constant for the mime type "application/vnd.xmpie.dpkg".
	ApplicationVndXmpieDpkg = "application/vnd.xmpie.dpkg"

	// ApplicationVndXmpiePlan is the constant for the mime type "application/vnd.xmpie.plan".
	ApplicationVndXmpiePlan = "application/vnd.xmpie.plan"

	// ApplicationVndXmpiePpkg is the constant for the mime type "application/vnd.xmpie.ppkg".
	ApplicationVndXmpiePpkg = "application/vnd.xmpie.ppkg"

	// ApplicationVndXmpieXlim is the constant for the mime type "application/vnd.xmpie.xlim".
	ApplicationVndXmpieXlim = "application/vnd.xmpie.xlim"

	// ApplicationVndYamahaHvDic is the constant for the mime type "application/vnd.yamaha.hv-dic".
	ApplicationVndYamahaHvDic = "application/vnd.yamaha.hv-dic"

	// ApplicationVndYamahaHvScript is the constant for the mime type "application/vnd.yamaha.hv-script".
	ApplicationVndYamahaHvScript = "application/vnd.yamaha.hv-script"

	// ApplicationVndYamahaHvVoice is the constant for the mime type "application/vnd.yamaha.hv-voice".
	ApplicationVndYamahaHvVoice = "application/vnd.yamaha.hv-voice"

	// ApplicationVndYamahaOpenscoreformatOsfpvgXml is the constant for the mime type "application/vnd.yamaha.openscoreformat.osfpvg+xml".
	ApplicationVndYamahaOpenscoreformatOsfpvgXml = "application/vnd.yamaha.openscoreformat.osfpvg+xml"

	// ApplicationVndYamahaOpenscoreformat is the constant for the mime type "application/vnd.yamaha.openscoreformat".
	ApplicationVndYamahaOpenscoreformat = "application/vnd.yamaha.openscoreformat"

	// ApplicationVndYamahaRemoteSetup is the constant for the mime type "application/vnd.yamaha.remote-setup".
	ApplicationVndYamahaRemoteSetup = "application/vnd.yamaha.remote-setup"

	// ApplicationVndYamahaSmafAudio is the constant for the mime type "application/vnd.yamaha.smaf-audio".
	ApplicationVndYamahaSmafAudio = "application/vnd.yamaha.smaf-audio"

	// ApplicationVndYamahaSmafPhrase is the constant for the mime type "application/vnd.yamaha.smaf-phrase".
	ApplicationVndYamahaSmafPhrase = "application/vnd.yamaha.smaf-phrase"

	// ApplicationVndYamahaThroughNgn is the constant for the mime type "application/vnd.yamaha.through-ngn".
	ApplicationVndYamahaThroughNgn = "application/vnd.yamaha.through-ngn"

	// ApplicationVndYamahaTunnelUdpencap is the constant for the mime type "application/vnd.yamaha.tunnel-udpencap".
	ApplicationVndYamahaTunnelUdpencap = "application/vnd.yamaha.tunnel-udpencap"

	// ApplicationVndYaoweme is the constant for the mime type "application/vnd.yaoweme".
	ApplicationVndYaoweme = "application/vnd.yaoweme"

	// ApplicationVndYellowriverCustomMenu is the constant for the mime type "application/vnd.yellowriver-custom-menu".
	ApplicationVndYellowriverCustomMenu = "application/vnd.yellowriver-custom-menu"

	// ApplicationVndYoutubeYt is the constant for the mime type "application/vnd.youtube.yt - OBSOLETED in favor of video/vnd.youtube.yt".
	ApplicationVndYoutubeYt = "application/vnd.youtube.yt"

	// ApplicationVndZul is the constant for the mime type "application/vnd.zul".
	ApplicationVndZul = "application/vnd.zul"

	// ApplicationVndZzazzDeckXml is the constant for the mime type "application/vnd.zzazz.deck+xml".
	ApplicationVndZzazzDeckXml = "application/vnd.zzazz.deck+xml"

	// ApplicationVoicexmlXml is the constant for the mime type "application/voicexml+xml".
	ApplicationVoicexmlXml = "application/voicexml+xml"

	// ApplicationVoucherCmsJson is the constant for the mime type "application/voucher-cms+json".
	ApplicationVoucherCmsJson = "application/voucher-cms+json"

	// ApplicationVqRtcpxr is the constant for the mime type "application/vq-rtcpxr".
	ApplicationVqRtcpxr = "application/vq-rtcpxr"

	// ApplicationWatcherinfoXml is the constant for the mime type "application/watcherinfo+xml".
	ApplicationWatcherinfoXml = "application/watcherinfo+xml"

	// ApplicationWebpushOptionsJson is the constant for the mime type "application/webpush-options+json".
	ApplicationWebpushOptionsJson = "application/webpush-options+json"

	// ApplicationWhoisppQuery is the constant for the mime type "application/whoispp-query".
	ApplicationWhoisppQuery = "application/whoispp-query"

	// ApplicationWhoisppResponse is the constant for the mime type "application/whoispp-response".
	ApplicationWhoisppResponse = "application/whoispp-response"

	// ApplicationWidget is the constant for the mime type "application/widget".
	ApplicationWidget = "application/widget"

	// ApplicationWita is the constant for the mime type "application/wita".
	ApplicationWita = "application/wita"

	// ApplicationWordperfect51 is the constant for the mime type "application/wordperfect5.1".
	ApplicationWordperfect51 = "application/wordperfect5.1"

	// ApplicationWsdlXml is the constant for the mime type "application/wsdl+xml".
	ApplicationWsdlXml = "application/wsdl+xml"

	// ApplicationWspolicyXml is the constant for the mime type "application/wspolicy+xml".
	ApplicationWspolicyXml = "application/wspolicy+xml"

	// ApplicationXWwwFormUrlencoded is the constant for the mime type "application/x-www-form-urlencoded".
	ApplicationXWwwFormUrlencoded = "application/x-www-form-urlencoded"

	// ApplicationX400Bp is the constant for the mime type "application/x400-bp".
	ApplicationX400Bp = "application/x400-bp"

	// ApplicationXacmlXml is the constant for the mime type "application/xacml+xml".
	ApplicationXacmlXml = "application/xacml+xml"

	// ApplicationXcapAttXml is the constant for the mime type "application/xcap-att+xml".
	ApplicationXcapAttXml = "application/xcap-att+xml"

	// ApplicationXcapCapsXml is the constant for the mime type "application/xcap-caps+xml".
	ApplicationXcapCapsXml = "application/xcap-caps+xml"

	// ApplicationXcapDiffXml is the constant for the mime type "application/xcap-diff+xml".
	ApplicationXcapDiffXml = "application/xcap-diff+xml"

	// ApplicationXcapElXml is the constant for the mime type "application/xcap-el+xml".
	ApplicationXcapElXml = "application/xcap-el+xml"

	// ApplicationXcapErrorXml is the constant for the mime type "application/xcap-error+xml".
	ApplicationXcapErrorXml = "application/xcap-error+xml"

	// ApplicationXcapNsXml is the constant for the mime type "application/xcap-ns+xml".
	ApplicationXcapNsXml = "application/xcap-ns+xml"

	// ApplicationXconConferenceInfoDiffXml is the constant for the mime type "application/xcon-conference-info-diff+xml".
	ApplicationXconConferenceInfoDiffXml = "application/xcon-conference-info-diff+xml"

	// ApplicationXconConferenceInfoXml is the constant for the mime type "application/xcon-conference-info+xml".
	ApplicationXconConferenceInfoXml = "application/xcon-conference-info+xml"

	// ApplicationXencXml is the constant for the mime type "application/xenc+xml".
	ApplicationXencXml = "application/xenc+xml"

	// ApplicationXhtmlXml is the constant for the mime type "application/xhtml+xml".
	ApplicationXhtmlXml = "application/xhtml+xml"

	// ApplicationXliffXml is the constant for the mime type "application/xliff+xml".
	ApplicationXliffXml = "application/xliff+xml"

	// ApplicationXml is the constant for the mime type "application/xml".
	ApplicationXml = "application/xml"

	// ApplicationXmlDtd is the constant for the mime type "application/xml-dtd".
	ApplicationXmlDtd = "application/xml-dtd"

	// ApplicationXmlExternalParsedEntity is the constant for the mime type "application/xml-external-parsed-entity".
	ApplicationXmlExternalParsedEntity = "application/xml-external-parsed-entity"

	// ApplicationXmlPatchXml is the constant for the mime type "application/xml-patch+xml".
	ApplicationXmlPatchXml = "application/xml-patch+xml"

	// ApplicationXmppXml is the constant for the mime type "application/xmpp+xml".
	ApplicationXmppXml = "application/xmpp+xml"

	// ApplicationXopXml is the constant for the mime type "application/xop+xml".
	ApplicationXopXml = "application/xop+xml"

	// ApplicationXsltXml is the constant for the mime type "application/xslt+xml".
	ApplicationXsltXml = "application/xslt+xml"

	// ApplicationXvXml is the constant for the mime type "application/xv+xml".
	ApplicationXvXml = "application/xv+xml"

	// ApplicationYang is the constant for the mime type "application/yang".
	ApplicationYang = "application/yang"

	// ApplicationYangDataJson is the constant for the mime type "application/yang-data+json".
	ApplicationYangDataJson = "application/yang-data+json"

	// ApplicationYangDataXml is the constant for the mime type "application/yang-data+xml".
	ApplicationYangDataXml = "application/yang-data+xml"

	// ApplicationYangPatchJson is the constant for the mime type "application/yang-patch+json".
	ApplicationYangPatchJson = "application/yang-patch+json"

	// ApplicationYangPatchXml is the constant for the mime type "application/yang-patch+xml".
	ApplicationYangPatchXml = "application/yang-patch+xml"

	// ApplicationYinXml is the constant for the mime type "application/yin+xml".
	ApplicationYinXml = "application/yin+xml"

	// ApplicationZip is the constant for the mime type "application/zip".
	ApplicationZip = "application/zip"

	// ApplicationZlib is the constant for the mime type "application/zlib".
	ApplicationZlib = "application/zlib"

	// ApplicationZstd is the constant for the mime type "application/zstd".
	ApplicationZstd = "application/zstd"
)

const (
	// Audio1dInterleavedParityfec is the constant for the mime type "audio/1d-interleaved-parityfec".
	Audio1dInterleavedParityfec = "audio/1d-interleaved-parityfec"

	// Audio32kadpcm is the constant for the mime type "audio/32kadpcm".
	Audio32kadpcm = "audio/32kadpcm"

	// Audio3gpp is the constant for the mime type "audio/3gpp".
	Audio3gpp = "audio/3gpp"

	// Audio3gpp2 is the constant for the mime type "audio/3gpp2".
	Audio3gpp2 = "audio/3gpp2"

	// AudioAac is the constant for the mime type "audio/aac".
	AudioAac = "audio/aac"

	// AudioAc3 is the constant for the mime type "audio/ac3".
	AudioAc3 = "audio/ac3"

	// AudioAMR is the constant for the mime type "audio/AMR".
	AudioAMR = "audio/AMR"

	// AudioAMRWB is the constant for the mime type "audio/AMR-WB".
	AudioAMRWB = "audio/AMR-WB"

	// AudioAmrWb is the constant for the mime type "audio/amr-wb+".
	AudioAmrWb = "audio/amr-wb+"

	// AudioAptx is the constant for the mime type "audio/aptx".
	AudioAptx = "audio/aptx"

	// AudioAsc is the constant for the mime type "audio/asc".
	AudioAsc = "audio/asc"

	// AudioATRACADVANCEDLOSSLESS is the constant for the mime type "audio/ATRAC-ADVANCED-LOSSLESS".
	AudioATRACADVANCEDLOSSLESS = "audio/ATRAC-ADVANCED-LOSSLESS"

	// AudioATRACX is the constant for the mime type "audio/ATRAC-X".
	AudioATRACX = "audio/ATRAC-X"

	// AudioATRAC3 is the constant for the mime type "audio/ATRAC3".
	AudioATRAC3 = "audio/ATRAC3"

	// AudioBasic is the constant for the mime type "audio/basic".
	AudioBasic = "audio/basic"

	// AudioBV16 is the constant for the mime type "audio/BV16".
	AudioBV16 = "audio/BV16"

	// AudioBV32 is the constant for the mime type "audio/BV32".
	AudioBV32 = "audio/BV32"

	// AudioClearmode is the constant for the mime type "audio/clearmode".
	AudioClearmode = "audio/clearmode"

	// AudioCN is the constant for the mime type "audio/CN".
	AudioCN = "audio/CN"

	// AudioDAT12 is the constant for the mime type "audio/DAT12".
	AudioDAT12 = "audio/DAT12"

	// AudioDls is the constant for the mime type "audio/dls".
	AudioDls = "audio/dls"

	// AudioDsrEs201108 is the constant for the mime type "audio/dsr-es201108".
	AudioDsrEs201108 = "audio/dsr-es201108"

	// AudioDsrEs202050 is the constant for the mime type "audio/dsr-es202050".
	AudioDsrEs202050 = "audio/dsr-es202050"

	// AudioDsrEs202211 is the constant for the mime type "audio/dsr-es202211".
	AudioDsrEs202211 = "audio/dsr-es202211"

	// AudioDsrEs202212 is the constant for the mime type "audio/dsr-es202212".
	AudioDsrEs202212 = "audio/dsr-es202212"

	// AudioDV is the constant for the mime type "audio/DV".
	AudioDV = "audio/DV"

	// AudioDVI4 is the constant for the mime type "audio/DVI4".
	AudioDVI4 = "audio/DVI4"

	// AudioEac3 is the constant for the mime type "audio/eac3".
	AudioEac3 = "audio/eac3"

	// AudioEncaprtp is the constant for the mime type "audio/encaprtp".
	AudioEncaprtp = "audio/encaprtp"

	// AudioEVRC is the constant for the mime type "audio/EVRC".
	AudioEVRC = "audio/EVRC"

	// AudioEVRCQCP is the constant for the mime type "audio/EVRC-QCP".
	AudioEVRCQCP = "audio/EVRC-QCP"

	// AudioEVRC0 is the constant for the mime type "audio/EVRC0".
	AudioEVRC0 = "audio/EVRC0"

	// AudioEVRC1 is the constant for the mime type "audio/EVRC1".
	AudioEVRC1 = "audio/EVRC1"

	// AudioEVRCB is the constant for the mime type "audio/EVRCB".
	AudioEVRCB = "audio/EVRCB"

	// AudioEVRCB0 is the constant for the mime type "audio/EVRCB0".
	AudioEVRCB0 = "audio/EVRCB0"

	// AudioEVRCB1 is the constant for the mime type "audio/EVRCB1".
	AudioEVRCB1 = "audio/EVRCB1"

	// AudioEVRCNW is the constant for the mime type "audio/EVRCNW".
	AudioEVRCNW = "audio/EVRCNW"

	// AudioEVRCNW0 is the constant for the mime type "audio/EVRCNW0".
	AudioEVRCNW0 = "audio/EVRCNW0"

	// AudioEVRCNW1 is the constant for the mime type "audio/EVRCNW1".
	AudioEVRCNW1 = "audio/EVRCNW1"

	// AudioEVRCWB is the constant for the mime type "audio/EVRCWB".
	AudioEVRCWB = "audio/EVRCWB"

	// AudioEVRCWB0 is the constant for the mime type "audio/EVRCWB0".
	AudioEVRCWB0 = "audio/EVRCWB0"

	// AudioEVRCWB1 is the constant for the mime type "audio/EVRCWB1".
	AudioEVRCWB1 = "audio/EVRCWB1"

	// AudioEVS is the constant for the mime type "audio/EVS".
	AudioEVS = "audio/EVS"

	// AudioExample is the constant for the mime type "audio/example".
	AudioExample = "audio/example"

	// AudioFlexfec is the constant for the mime type "audio/flexfec".
	AudioFlexfec = "audio/flexfec"

	// AudioFwdred is the constant for the mime type "audio/fwdred".
	AudioFwdred = "audio/fwdred"

	// AudioG7110 is the constant for the mime type "audio/G711-0".
	AudioG7110 = "audio/G711-0"

	// AudioG719 is the constant for the mime type "audio/G719".
	AudioG719 = "audio/G719"

	// AudioG7221 is the constant for the mime type "audio/G7221".
	AudioG7221 = "audio/G7221"

	// AudioG722 is the constant for the mime type "audio/G722".
	AudioG722 = "audio/G722"

	// AudioG723 is the constant for the mime type "audio/G723".
	AudioG723 = "audio/G723"

	// AudioG72616 is the constant for the mime type "audio/G726-16".
	AudioG72616 = "audio/G726-16"

	// AudioG72624 is the constant for the mime type "audio/G726-24".
	AudioG72624 = "audio/G726-24"

	// AudioG72632 is the constant for the mime type "audio/G726-32".
	AudioG72632 = "audio/G726-32"

	// AudioG72640 is the constant for the mime type "audio/G726-40".
	AudioG72640 = "audio/G726-40"

	// AudioG728 is the constant for the mime type "audio/G728".
	AudioG728 = "audio/G728"

	// AudioG729 is the constant for the mime type "audio/G729".
	AudioG729 = "audio/G729"

	// AudioG7291 is the constant for the mime type "audio/G7291".
	AudioG7291 = "audio/G7291"

	// AudioG729D is the constant for the mime type "audio/G729D".
	AudioG729D = "audio/G729D"

	// AudioG729E is the constant for the mime type "audio/G729E".
	AudioG729E = "audio/G729E"

	// AudioGSM is the constant for the mime type "audio/GSM".
	AudioGSM = "audio/GSM"

	// AudioGSMEFR is the constant for the mime type "audio/GSM-EFR".
	AudioGSMEFR = "audio/GSM-EFR"

	// AudioGSMHR08 is the constant for the mime type "audio/GSM-HR-08".
	AudioGSMHR08 = "audio/GSM-HR-08"

	// AudioILBC is the constant for the mime type "audio/iLBC".
	AudioILBC = "audio/iLBC"

	// AudioIpMrV25 is the constant for the mime type "audio/ip-mr_v2.5".
	AudioIpMrV25 = "audio/ip-mr_v2.5"

	// AudioL8 is the constant for the mime type "audio/L8".
	AudioL8 = "audio/L8"

	// AudioL16 is the constant for the mime type "audio/L16".
	AudioL16 = "audio/L16"

	// AudioL20 is the constant for the mime type "audio/L20".
	AudioL20 = "audio/L20"

	// AudioL24 is the constant for the mime type "audio/L24".
	AudioL24 = "audio/L24"

	// AudioLPC is the constant for the mime type "audio/LPC".
	AudioLPC = "audio/LPC"

	// AudioMELP is the constant for the mime type "audio/MELP".
	AudioMELP = "audio/MELP"

	// AudioMELP600 is the constant for the mime type "audio/MELP600".
	AudioMELP600 = "audio/MELP600"

	// AudioMELP1200 is the constant for the mime type "audio/MELP1200".
	AudioMELP1200 = "audio/MELP1200"

	// AudioMELP2400 is the constant for the mime type "audio/MELP2400".
	AudioMELP2400 = "audio/MELP2400"

	// AudioMobileXmf is the constant for the mime type "audio/mobile-xmf".
	AudioMobileXmf = "audio/mobile-xmf"

	// AudioMPA is the constant for the mime type "audio/MPA".
	AudioMPA = "audio/MPA"

	// AudioMp4 is the constant for the mime type "audio/mp4".
	AudioMp4 = "audio/mp4"

	// AudioMP4ALATM is the constant for the mime type "audio/MP4A-LATM".
	AudioMP4ALATM = "audio/MP4A-LATM"

	// AudioMpaRobust is the constant for the mime type "audio/mpa-robust".
	AudioMpaRobust = "audio/mpa-robust"

	// AudioMpeg is the constant for the mime type "audio/mpeg".
	AudioMpeg = "audio/mpeg"

	// AudioMpeg4Generic is the constant for the mime type "audio/mpeg4-generic".
	AudioMpeg4Generic = "audio/mpeg4-generic"

	// AudioOgg is the constant for the mime type "audio/ogg".
	AudioOgg = "audio/ogg"

	// AudioOpus is the constant for the mime type "audio/opus".
	AudioOpus = "audio/opus"

	// AudioParityfec is the constant for the mime type "audio/parityfec".
	AudioParityfec = "audio/parityfec"

	// AudioPCMA is the constant for the mime type "audio/PCMA".
	AudioPCMA = "audio/PCMA"

	// AudioPCMAWB is the constant for the mime type "audio/PCMA-WB".
	AudioPCMAWB = "audio/PCMA-WB"

	// AudioPCMU is the constant for the mime type "audio/PCMU".
	AudioPCMU = "audio/PCMU"

	// AudioPCMUWB is the constant for the mime type "audio/PCMU-WB".
	AudioPCMUWB = "audio/PCMU-WB"

	// AudioPrsSid is the constant for the mime type "audio/prs.sid".
	AudioPrsSid = "audio/prs.sid"

	// AudioQCELP is the constant for the mime type "audio/QCELP".
	AudioQCELP = "audio/QCELP"

	// AudioRaptorfec is the constant for the mime type "audio/raptorfec".
	AudioRaptorfec = "audio/raptorfec"

	// AudioRED is the constant for the mime type "audio/RED".
	AudioRED = "audio/RED"

	// AudioRtpEncAescm128 is the constant for the mime type "audio/rtp-enc-aescm128".
	AudioRtpEncAescm128 = "audio/rtp-enc-aescm128"

	// AudioRtploopback is the constant for the mime type "audio/rtploopback".
	AudioRtploopback = "audio/rtploopback"

	// AudioRtpMidi is the constant for the mime type "audio/rtp-midi".
	AudioRtpMidi = "audio/rtp-midi"

	// AudioRtx is the constant for the mime type "audio/rtx".
	AudioRtx = "audio/rtx"

	// AudioSMV is the constant for the mime type "audio/SMV".
	AudioSMV = "audio/SMV"

	// AudioSMV0 is the constant for the mime type "audio/SMV0".
	AudioSMV0 = "audio/SMV0"

	// AudioSMVQCP is the constant for the mime type "audio/SMV-QCP".
	AudioSMVQCP = "audio/SMV-QCP"

	// AudioSpMidi is the constant for the mime type "audio/sp-midi".
	AudioSpMidi = "audio/sp-midi"

	// AudioSpeex is the constant for the mime type "audio/speex".
	AudioSpeex = "audio/speex"

	// AudioT140C is the constant for the mime type "audio/t140c".
	AudioT140C = "audio/t140c"

	// AudioT38 is the constant for the mime type "audio/t38".
	AudioT38 = "audio/t38"

	// AudioTelephoneEvent is the constant for the mime type "audio/telephone-event".
	AudioTelephoneEvent = "audio/telephone-event"

	// AudioTETRAACELP is the constant for the mime type "audio/TETRA_ACELP".
	AudioTETRAACELP = "audio/TETRA_ACELP"

	// AudioTone is the constant for the mime type "audio/tone".
	AudioTone = "audio/tone"

	// AudioUEMCLIP is the constant for the mime type "audio/UEMCLIP".
	AudioUEMCLIP = "audio/UEMCLIP"

	// AudioUlpfec is the constant for the mime type "audio/ulpfec".
	AudioUlpfec = "audio/ulpfec"

	// AudioUsac is the constant for the mime type "audio/usac".
	AudioUsac = "audio/usac"

	// AudioVDVI is the constant for the mime type "audio/VDVI".
	AudioVDVI = "audio/VDVI"

	// AudioVMRWB is the constant for the mime type "audio/VMR-WB".
	AudioVMRWB = "audio/VMR-WB"

	// AudioVnd3gppIufp is the constant for the mime type "audio/vnd.3gpp.iufp".
	AudioVnd3gppIufp = "audio/vnd.3gpp.iufp"

	// AudioVnd4SB is the constant for the mime type "audio/vnd.4SB".
	AudioVnd4SB = "audio/vnd.4SB"

	// AudioVndAudiokoz is the constant for the mime type "audio/vnd.audiokoz".
	AudioVndAudiokoz = "audio/vnd.audiokoz"

	// AudioVndCELP is the constant for the mime type "audio/vnd.CELP".
	AudioVndCELP = "audio/vnd.CELP"

	// AudioVndCiscoNse is the constant for the mime type "audio/vnd.cisco.nse".
	AudioVndCiscoNse = "audio/vnd.cisco.nse"

	// AudioVndCmlesRadioEvents is the constant for the mime type "audio/vnd.cmles.radio-events".
	AudioVndCmlesRadioEvents = "audio/vnd.cmles.radio-events"

	// AudioVndCnsAnp1 is the constant for the mime type "audio/vnd.cns.anp1".
	AudioVndCnsAnp1 = "audio/vnd.cns.anp1"

	// AudioVndCnsInf1 is the constant for the mime type "audio/vnd.cns.inf1".
	AudioVndCnsInf1 = "audio/vnd.cns.inf1"

	// AudioVndDeceAudio is the constant for the mime type "audio/vnd.dece.audio".
	AudioVndDeceAudio = "audio/vnd.dece.audio"

	// AudioVndDigitalWinds is the constant for the mime type "audio/vnd.digital-winds".
	AudioVndDigitalWinds = "audio/vnd.digital-winds"

	// AudioVndDlnaAdts is the constant for the mime type "audio/vnd.dlna.adts".
	AudioVndDlnaAdts = "audio/vnd.dlna.adts"

	// AudioVndDolbyHeaac1 is the constant for the mime type "audio/vnd.dolby.heaac.1".
	AudioVndDolbyHeaac1 = "audio/vnd.dolby.heaac.1"

	// AudioVndDolbyHeaac2 is the constant for the mime type "audio/vnd.dolby.heaac.2".
	AudioVndDolbyHeaac2 = "audio/vnd.dolby.heaac.2"

	// AudioVndDolbyMlp is the constant for the mime type "audio/vnd.dolby.mlp".
	AudioVndDolbyMlp = "audio/vnd.dolby.mlp"

	// AudioVndDolbyMps is the constant for the mime type "audio/vnd.dolby.mps".
	AudioVndDolbyMps = "audio/vnd.dolby.mps"

	// AudioVndDolbyPl2 is the constant for the mime type "audio/vnd.dolby.pl2".
	AudioVndDolbyPl2 = "audio/vnd.dolby.pl2"

	// AudioVndDolbyPl2X is the constant for the mime type "audio/vnd.dolby.pl2x".
	AudioVndDolbyPl2X = "audio/vnd.dolby.pl2x"

	// AudioVndDolbyPl2Z is the constant for the mime type "audio/vnd.dolby.pl2z".
	AudioVndDolbyPl2Z = "audio/vnd.dolby.pl2z"

	// AudioVndDolbyPulse1 is the constant for the mime type "audio/vnd.dolby.pulse.1".
	AudioVndDolbyPulse1 = "audio/vnd.dolby.pulse.1"

	// AudioVndDra is the constant for the mime type "audio/vnd.dra".
	AudioVndDra = "audio/vnd.dra"

	// AudioVndDts is the constant for the mime type "audio/vnd.dts".
	AudioVndDts = "audio/vnd.dts"

	// AudioVndDtsHd is the constant for the mime type "audio/vnd.dts.hd".
	AudioVndDtsHd = "audio/vnd.dts.hd"

	// AudioVndDtsUhd is the constant for the mime type "audio/vnd.dts.uhd".
	AudioVndDtsUhd = "audio/vnd.dts.uhd"

	// AudioVndDvbFile is the constant for the mime type "audio/vnd.dvb.file".
	AudioVndDvbFile = "audio/vnd.dvb.file"

	// AudioVndEveradPlj is the constant for the mime type "audio/vnd.everad.plj".
	AudioVndEveradPlj = "audio/vnd.everad.plj"

	// AudioVndHnsAudio is the constant for the mime type "audio/vnd.hns.audio".
	AudioVndHnsAudio = "audio/vnd.hns.audio"

	// AudioVndLucentVoice is the constant for the mime type "audio/vnd.lucent.voice".
	AudioVndLucentVoice = "audio/vnd.lucent.voice"

	// AudioVndMsPlayreadyMediaPya is the constant for the mime type "audio/vnd.ms-playready.media.pya".
	AudioVndMsPlayreadyMediaPya = "audio/vnd.ms-playready.media.pya"

	// AudioVndNokiaMobileXmf is the constant for the mime type "audio/vnd.nokia.mobile-xmf".
	AudioVndNokiaMobileXmf = "audio/vnd.nokia.mobile-xmf"

	// AudioVndNortelVbk is the constant for the mime type "audio/vnd.nortel.vbk".
	AudioVndNortelVbk = "audio/vnd.nortel.vbk"

	// AudioVndNueraEcelp4800 is the constant for the mime type "audio/vnd.nuera.ecelp4800".
	AudioVndNueraEcelp4800 = "audio/vnd.nuera.ecelp4800"

	// AudioVndNueraEcelp7470 is the constant for the mime type "audio/vnd.nuera.ecelp7470".
	AudioVndNueraEcelp7470 = "audio/vnd.nuera.ecelp7470"

	// AudioVndNueraEcelp9600 is the constant for the mime type "audio/vnd.nuera.ecelp9600".
	AudioVndNueraEcelp9600 = "audio/vnd.nuera.ecelp9600"

	// AudioVndOctelSbc is the constant for the mime type "audio/vnd.octel.sbc".
	AudioVndOctelSbc = "audio/vnd.octel.sbc"

	// AudioVndPresonusMultitrack is the constant for the mime type "audio/vnd.presonus.multitrack".
	AudioVndPresonusMultitrack = "audio/vnd.presonus.multitrack"

	// AudioVndQcelp is the constant for the mime type "audio/vnd.qcelp - DEPRECATED in favor of audio/qcelp".
	AudioVndQcelp = "audio/vnd.qcelp"

	// AudioVndRhetorex32kadpcm is the constant for the mime type "audio/vnd.rhetorex.32kadpcm".
	AudioVndRhetorex32kadpcm = "audio/vnd.rhetorex.32kadpcm"

	// AudioVndRip is the constant for the mime type "audio/vnd.rip".
	AudioVndRip = "audio/vnd.rip"

	// AudioVndSealedmediaSoftsealMpeg is the constant for the mime type "audio/vnd.sealedmedia.softseal.mpeg".
	AudioVndSealedmediaSoftsealMpeg = "audio/vnd.sealedmedia.softseal.mpeg"

	// AudioVndVmxCvsd is the constant for the mime type "audio/vnd.vmx.cvsd".
	AudioVndVmxCvsd = "audio/vnd.vmx.cvsd"

	// AudioVorbis is the constant for the mime type "audio/vorbis".
	AudioVorbis = "audio/vorbis"

	// AudioVorbisConfig is the constant for the mime type "audio/vorbis-config".
	AudioVorbisConfig = "audio/vorbis-config"
)

const (
	// FontCollection is the constant for the mime type "font/collection".
	FontCollection = "font/collection"

	// FontOtf is the constant for the mime type "font/otf".
	FontOtf = "font/otf"

	// FontSfnt is the constant for the mime type "font/sfnt".
	FontSfnt = "font/sfnt"

	// FontTtf is the constant for the mime type "font/ttf".
	FontTtf = "font/ttf"

	// FontWoff is the constant for the mime type "font/woff".
	FontWoff = "font/woff"

	// FontWoff2 is the constant for the mime type "font/woff2".
	FontWoff2 = "font/woff2"
)

const ()

const (
	// ImageAces is the constant for the mime type "image/aces".
	ImageAces = "image/aces"

	// ImageAvci is the constant for the mime type "image/avci".
	ImageAvci = "image/avci"

	// ImageAvcs is the constant for the mime type "image/avcs".
	ImageAvcs = "image/avcs"

	// ImageBmp is the constant for the mime type "image/bmp".
	ImageBmp = "image/bmp"

	// ImageCgm is the constant for the mime type "image/cgm".
	ImageCgm = "image/cgm"

	// ImageDicomRle is the constant for the mime type "image/dicom-rle".
	ImageDicomRle = "image/dicom-rle"

	// ImageEmf is the constant for the mime type "image/emf".
	ImageEmf = "image/emf"

	// ImageExample is the constant for the mime type "image/example".
	ImageExample = "image/example"

	// ImageFits is the constant for the mime type "image/fits".
	ImageFits = "image/fits"

	// ImageG3Fax is the constant for the mime type "image/g3fax".
	ImageG3Fax = "image/g3fax"

	// ImageGif is the constant for the mime type "image/gif".
	ImageGif = "image/gif"

	// ImageHeic is the constant for the mime type "image/heic".
	ImageHeic = "image/heic"

	// ImageHeicSequence is the constant for the mime type "image/heic-sequence".
	ImageHeicSequence = "image/heic-sequence"

	// ImageHeif is the constant for the mime type "image/heif".
	ImageHeif = "image/heif"

	// ImageHeifSequence is the constant for the mime type "image/heif-sequence".
	ImageHeifSequence = "image/heif-sequence"

	// ImageHej2K is the constant for the mime type "image/hej2k".
	ImageHej2K = "image/hej2k"

	// ImageHsj2 is the constant for the mime type "image/hsj2".
	ImageHsj2 = "image/hsj2"

	// ImageIef is the constant for the mime type "image/ief".
	ImageIef = "image/ief"

	// ImageJls is the constant for the mime type "image/jls".
	ImageJls = "image/jls"

	// ImageJp2 is the constant for the mime type "image/jp2".
	ImageJp2 = "image/jp2"

	// ImageJpeg is the constant for the mime type "image/jpeg".
	ImageJpeg = "image/jpeg"

	// ImageJph is the constant for the mime type "image/jph".
	ImageJph = "image/jph"

	// ImageJphc is the constant for the mime type "image/jphc".
	ImageJphc = "image/jphc"

	// ImageJpm is the constant for the mime type "image/jpm".
	ImageJpm = "image/jpm"

	// ImageJpx is the constant for the mime type "image/jpx".
	ImageJpx = "image/jpx"

	// ImageJxr is the constant for the mime type "image/jxr".
	ImageJxr = "image/jxr"

	// ImageJxrA is the constant for the mime type "image/jxrA".
	ImageJxrA = "image/jxrA"

	// ImageJxrS is the constant for the mime type "image/jxrS".
	ImageJxrS = "image/jxrS"

	// ImageJxs is the constant for the mime type "image/jxs".
	ImageJxs = "image/jxs"

	// ImageJxsc is the constant for the mime type "image/jxsc".
	ImageJxsc = "image/jxsc"

	// ImageJxsi is the constant for the mime type "image/jxsi".
	ImageJxsi = "image/jxsi"

	// ImageJxss is the constant for the mime type "image/jxss".
	ImageJxss = "image/jxss"

	// ImageKtx is the constant for the mime type "image/ktx".
	ImageKtx = "image/ktx"

	// ImageNaplps is the constant for the mime type "image/naplps".
	ImageNaplps = "image/naplps"

	// ImagePng is the constant for the mime type "image/png".
	ImagePng = "image/png"

	// ImagePrsBtif is the constant for the mime type "image/prs.btif".
	ImagePrsBtif = "image/prs.btif"

	// ImagePrsPti is the constant for the mime type "image/prs.pti".
	ImagePrsPti = "image/prs.pti"

	// ImagePwgRaster is the constant for the mime type "image/pwg-raster".
	ImagePwgRaster = "image/pwg-raster"

	// ImageSvgXml is the constant for the mime type "image/svg+xml".
	ImageSvgXml = "image/svg+xml"

	// ImageT38 is the constant for the mime type "image/t38".
	ImageT38 = "image/t38"

	// ImageTiff is the constant for the mime type "image/tiff".
	ImageTiff = "image/tiff"

	// ImageTiffFx is the constant for the mime type "image/tiff-fx".
	ImageTiffFx = "image/tiff-fx"

	// ImageVndAdobePhotoshop is the constant for the mime type "image/vnd.adobe.photoshop".
	ImageVndAdobePhotoshop = "image/vnd.adobe.photoshop"

	// ImageVndAirzipAcceleratorAzv is the constant for the mime type "image/vnd.airzip.accelerator.azv".
	ImageVndAirzipAcceleratorAzv = "image/vnd.airzip.accelerator.azv"

	// ImageVndCnsInf2 is the constant for the mime type "image/vnd.cns.inf2".
	ImageVndCnsInf2 = "image/vnd.cns.inf2"

	// ImageVndDeceGraphic is the constant for the mime type "image/vnd.dece.graphic".
	ImageVndDeceGraphic = "image/vnd.dece.graphic"

	// ImageVndDjvu is the constant for the mime type "image/vnd.djvu".
	ImageVndDjvu = "image/vnd.djvu"

	// ImageVndDwg is the constant for the mime type "image/vnd.dwg".
	ImageVndDwg = "image/vnd.dwg"

	// ImageVndDxf is the constant for the mime type "image/vnd.dxf".
	ImageVndDxf = "image/vnd.dxf"

	// ImageVndDvbSubtitle is the constant for the mime type "image/vnd.dvb.subtitle".
	ImageVndDvbSubtitle = "image/vnd.dvb.subtitle"

	// ImageVndFastbidsheet is the constant for the mime type "image/vnd.fastbidsheet".
	ImageVndFastbidsheet = "image/vnd.fastbidsheet"

	// ImageVndFpx is the constant for the mime type "image/vnd.fpx".
	ImageVndFpx = "image/vnd.fpx"

	// ImageVndFst is the constant for the mime type "image/vnd.fst".
	ImageVndFst = "image/vnd.fst"

	// ImageVndFujixeroxEdmicsMmr is the constant for the mime type "image/vnd.fujixerox.edmics-mmr".
	ImageVndFujixeroxEdmicsMmr = "image/vnd.fujixerox.edmics-mmr"

	// ImageVndFujixeroxEdmicsRlc is the constant for the mime type "image/vnd.fujixerox.edmics-rlc".
	ImageVndFujixeroxEdmicsRlc = "image/vnd.fujixerox.edmics-rlc"

	// ImageVndGlobalgraphicsPgb is the constant for the mime type "image/vnd.globalgraphics.pgb".
	ImageVndGlobalgraphicsPgb = "image/vnd.globalgraphics.pgb"

	// ImageVndMicrosoftIcon is the constant for the mime type "image/vnd.microsoft.icon".
	ImageVndMicrosoftIcon = "image/vnd.microsoft.icon"

	// ImageVndMix is the constant for the mime type "image/vnd.mix".
	ImageVndMix = "image/vnd.mix"

	// ImageVndMsModi is the constant for the mime type "image/vnd.ms-modi".
	ImageVndMsModi = "image/vnd.ms-modi"

	// ImageVndMozillaApng is the constant for the mime type "image/vnd.mozilla.apng".
	ImageVndMozillaApng = "image/vnd.mozilla.apng"

	// ImageVndNetFpx is the constant for the mime type "image/vnd.net-fpx".
	ImageVndNetFpx = "image/vnd.net-fpx"

	// ImageVndRadiance is the constant for the mime type "image/vnd.radiance".
	ImageVndRadiance = "image/vnd.radiance"

	// ImageVndSealedPng is the constant for the mime type "image/vnd.sealed.png".
	ImageVndSealedPng = "image/vnd.sealed.png"

	// ImageVndSealedmediaSoftsealGif is the constant for the mime type "image/vnd.sealedmedia.softseal.gif".
	ImageVndSealedmediaSoftsealGif = "image/vnd.sealedmedia.softseal.gif"

	// ImageVndSealedmediaSoftsealJpg is the constant for the mime type "image/vnd.sealedmedia.softseal.jpg".
	ImageVndSealedmediaSoftsealJpg = "image/vnd.sealedmedia.softseal.jpg"

	// ImageVndSvf is the constant for the mime type "image/vnd.svf".
	ImageVndSvf = "image/vnd.svf"

	// ImageVndTencentTap is the constant for the mime type "image/vnd.tencent.tap".
	ImageVndTencentTap = "image/vnd.tencent.tap"

	// ImageVndValveSourceTexture is the constant for the mime type "image/vnd.valve.source.texture".
	ImageVndValveSourceTexture = "image/vnd.valve.source.texture"

	// ImageVndWapWbmp is the constant for the mime type "image/vnd.wap.wbmp".
	ImageVndWapWbmp = "image/vnd.wap.wbmp"

	// ImageVndXiff is the constant for the mime type "image/vnd.xiff".
	ImageVndXiff = "image/vnd.xiff"

	// ImageVndZbrushPcx is the constant for the mime type "image/vnd.zbrush.pcx".
	ImageVndZbrushPcx = "image/vnd.zbrush.pcx"

	// ImageWmf is the constant for the mime type "image/wmf".
	ImageWmf = "image/wmf"

	// ImageXEmf is the constant for the mime type "image/x-emf - DEPRECATED in favor of image/emf".
	ImageXEmf = "image/x-emf"

	// ImageXWmf is the constant for the mime type "image/x-wmf - DEPRECATED in favor of image/wmf".
	ImageXWmf = "image/x-wmf"
)

const (
	// MessageCPIM is the constant for the mime type "message/CPIM".
	MessageCPIM = "message/CPIM"

	// MessageDeliveryStatus is the constant for the mime type "message/delivery-status".
	MessageDeliveryStatus = "message/delivery-status"

	// MessageDispositionNotification is the constant for the mime type "message/disposition-notification".
	MessageDispositionNotification = "message/disposition-notification"

	// MessageExample is the constant for the mime type "message/example".
	MessageExample = "message/example"

	// MessageExternalBody is the constant for the mime type "message/external-body".
	MessageExternalBody = "message/external-body"

	// MessageFeedbackReport is the constant for the mime type "message/feedback-report".
	MessageFeedbackReport = "message/feedback-report"

	// MessageGlobal is the constant for the mime type "message/global".
	MessageGlobal = "message/global"

	// MessageGlobalDeliveryStatus is the constant for the mime type "message/global-delivery-status".
	MessageGlobalDeliveryStatus = "message/global-delivery-status"

	// MessageGlobalDispositionNotification is the constant for the mime type "message/global-disposition-notification".
	MessageGlobalDispositionNotification = "message/global-disposition-notification"

	// MessageGlobalHeaders is the constant for the mime type "message/global-headers".
	MessageGlobalHeaders = "message/global-headers"

	// MessageHttp is the constant for the mime type "message/http".
	MessageHttp = "message/http"

	// MessageImdnXml is the constant for the mime type "message/imdn+xml".
	MessageImdnXml = "message/imdn+xml"

	// MessageNews is the constant for the mime type "message/news - OBSOLETED by RFC5537".
	MessageNews = "message/news"

	// MessagePartial is the constant for the mime type "message/partial".
	MessagePartial = "message/partial"

	// MessageRfc822 is the constant for the mime type "message/rfc822".
	MessageRfc822 = "message/rfc822"

	// MessageSHttp is the constant for the mime type "message/s-http".
	MessageSHttp = "message/s-http"

	// MessageSip is the constant for the mime type "message/sip".
	MessageSip = "message/sip"

	// MessageSipfrag is the constant for the mime type "message/sipfrag".
	MessageSipfrag = "message/sipfrag"

	// MessageTrackingStatus is the constant for the mime type "message/tracking-status".
	MessageTrackingStatus = "message/tracking-status"

	// MessageVndSiSimp is the constant for the mime type "message/vnd.si.simp - OBSOLETED by request".
	MessageVndSiSimp = "message/vnd.si.simp"

	// MessageVndWfaWsc is the constant for the mime type "message/vnd.wfa.wsc".
	MessageVndWfaWsc = "message/vnd.wfa.wsc"
)

const (
	// Model3mf is the constant for the mime type "model/3mf".
	Model3mf = "model/3mf"

	// ModelExample is the constant for the mime type "model/example".
	ModelExample = "model/example"

	// ModelGltfBinary is the constant for the mime type "model/gltf-binary".
	ModelGltfBinary = "model/gltf-binary"

	// ModelGltfJson is the constant for the mime type "model/gltf+json".
	ModelGltfJson = "model/gltf+json"

	// ModelIges is the constant for the mime type "model/iges".
	ModelIges = "model/iges"

	// ModelMesh is the constant for the mime type "model/mesh".
	ModelMesh = "model/mesh"

	// ModelStl is the constant for the mime type "model/stl".
	ModelStl = "model/stl"

	// ModelVndColladaXml is the constant for the mime type "model/vnd.collada+xml".
	ModelVndColladaXml = "model/vnd.collada+xml"

	// ModelVndDwf is the constant for the mime type "model/vnd.dwf".
	ModelVndDwf = "model/vnd.dwf"

	// ModelVndFlatland3dml is the constant for the mime type "model/vnd.flatland.3dml".
	ModelVndFlatland3dml = "model/vnd.flatland.3dml"

	// ModelVndGdl is the constant for the mime type "model/vnd.gdl".
	ModelVndGdl = "model/vnd.gdl"

	// ModelVndGsGdl is the constant for the mime type "model/vnd.gs-gdl".
	ModelVndGsGdl = "model/vnd.gs-gdl"

	// ModelVndGtw is the constant for the mime type "model/vnd.gtw".
	ModelVndGtw = "model/vnd.gtw"

	// ModelVndMomlXml is the constant for the mime type "model/vnd.moml+xml".
	ModelVndMomlXml = "model/vnd.moml+xml"

	// ModelVndMts is the constant for the mime type "model/vnd.mts".
	ModelVndMts = "model/vnd.mts"

	// ModelVndOpengex is the constant for the mime type "model/vnd.opengex".
	ModelVndOpengex = "model/vnd.opengex"

	// ModelVndParasolidTransmitBinary is the constant for the mime type "model/vnd.parasolid.transmit.binary".
	ModelVndParasolidTransmitBinary = "model/vnd.parasolid.transmit.binary"

	// ModelVndParasolidTransmitText is the constant for the mime type "model/vnd.parasolid.transmit.text".
	ModelVndParasolidTransmitText = "model/vnd.parasolid.transmit.text"

	// ModelVndRosetteAnnotatedDataModel is the constant for the mime type "model/vnd.rosette.annotated-data-model".
	ModelVndRosetteAnnotatedDataModel = "model/vnd.rosette.annotated-data-model"

	// ModelVndUsdzZip is the constant for the mime type "model/vnd.usdz+zip".
	ModelVndUsdzZip = "model/vnd.usdz+zip"

	// ModelVndValveSourceCompiledMap is the constant for the mime type "model/vnd.valve.source.compiled-map".
	ModelVndValveSourceCompiledMap = "model/vnd.valve.source.compiled-map"

	// ModelVndVtu is the constant for the mime type "model/vnd.vtu".
	ModelVndVtu = "model/vnd.vtu"

	// ModelVrml is the constant for the mime type "model/vrml".
	ModelVrml = "model/vrml"

	// ModelX3DVrml is the constant for the mime type "model/x3d-vrml".
	ModelX3DVrml = "model/x3d-vrml"

	// ModelX3DFastinfoset is the constant for the mime type "model/x3d+fastinfoset".
	ModelX3DFastinfoset = "model/x3d+fastinfoset"

	// ModelX3DXml is the constant for the mime type "model/x3d+xml".
	ModelX3DXml = "model/x3d+xml"
)

const (
	// MultipartAlternative is the constant for the mime type "multipart/alternative".
	MultipartAlternative = "multipart/alternative"

	// MultipartAppledouble is the constant for the mime type "multipart/appledouble".
	MultipartAppledouble = "multipart/appledouble"

	// MultipartByteranges is the constant for the mime type "multipart/byteranges".
	MultipartByteranges = "multipart/byteranges"

	// MultipartDigest is the constant for the mime type "multipart/digest".
	MultipartDigest = "multipart/digest"

	// MultipartEncrypted is the constant for the mime type "multipart/encrypted".
	MultipartEncrypted = "multipart/encrypted"

	// MultipartExample is the constant for the mime type "multipart/example".
	MultipartExample = "multipart/example"

	// MultipartFormData is the constant for the mime type "multipart/form-data".
	MultipartFormData = "multipart/form-data"

	// MultipartHeaderSet is the constant for the mime type "multipart/header-set".
	MultipartHeaderSet = "multipart/header-set"

	// MultipartMixed is the constant for the mime type "multipart/mixed".
	MultipartMixed = "multipart/mixed"

	// MultipartMultilingual is the constant for the mime type "multipart/multilingual".
	MultipartMultilingual = "multipart/multilingual"

	// MultipartParallel is the constant for the mime type "multipart/parallel".
	MultipartParallel = "multipart/parallel"

	// MultipartRelated is the constant for the mime type "multipart/related".
	MultipartRelated = "multipart/related"

	// MultipartReport is the constant for the mime type "multipart/report".
	MultipartReport = "multipart/report"

	// MultipartSigned is the constant for the mime type "multipart/signed".
	MultipartSigned = "multipart/signed"

	// MultipartVndBintMedPlus is the constant for the mime type "multipart/vnd.bint.med-plus".
	MultipartVndBintMedPlus = "multipart/vnd.bint.med-plus"

	// MultipartVoiceMessage is the constant for the mime type "multipart/voice-message".
	MultipartVoiceMessage = "multipart/voice-message"

	// MultipartXMixedReplace is the constant for the mime type "multipart/x-mixed-replace".
	MultipartXMixedReplace = "multipart/x-mixed-replace"
)

const (
	// Text1dInterleavedParityfec is the constant for the mime type "text/1d-interleaved-parityfec".
	Text1dInterleavedParityfec = "text/1d-interleaved-parityfec"

	// TextCacheManifest is the constant for the mime type "text/cache-manifest".
	TextCacheManifest = "text/cache-manifest"

	// TextCalendar is the constant for the mime type "text/calendar".
	TextCalendar = "text/calendar"

	// TextCss is the constant for the mime type "text/css".
	TextCss = "text/css"

	// TextCsv is the constant for the mime type "text/csv".
	TextCsv = "text/csv"

	// TextCsvSchema is the constant for the mime type "text/csv-schema".
	TextCsvSchema = "text/csv-schema"

	// TextDirectory is the constant for the mime type "text/directory - DEPRECATED by RFC6350".
	TextDirectory = "text/directory"

	// TextDns is the constant for the mime type "text/dns".
	TextDns = "text/dns"

	// TextEcmascript is the constant for the mime type "text/ecmascript - OBSOLETED in favor of application/ecmascript".
	TextEcmascript = "text/ecmascript"

	// TextEncaprtp is the constant for the mime type "text/encaprtp".
	TextEncaprtp = "text/encaprtp"

	// TextEnriched is the constant for the mime type "text/enriched".
	TextEnriched = "text/enriched"

	// TextExample is the constant for the mime type "text/example".
	TextExample = "text/example"

	// TextFlexfec is the constant for the mime type "text/flexfec".
	TextFlexfec = "text/flexfec"

	// TextFwdred is the constant for the mime type "text/fwdred".
	TextFwdred = "text/fwdred"

	// TextGrammarRefList is the constant for the mime type "text/grammar-ref-list".
	TextGrammarRefList = "text/grammar-ref-list"

	// TextHtml is the constant for the mime type "text/html".
	TextHtml = "text/html"

	// TextJavascript is the constant for the mime type "text/javascript - OBSOLETED in favor of application/javascript".
	TextJavascript = "text/javascript"

	// TextJcrCnd is the constant for the mime type "text/jcr-cnd".
	TextJcrCnd = "text/jcr-cnd"

	// TextMarkdown is the constant for the mime type "text/markdown".
	TextMarkdown = "text/markdown"

	// TextMizar is the constant for the mime type "text/mizar".
	TextMizar = "text/mizar"

	// TextN3 is the constant for the mime type "text/n3".
	TextN3 = "text/n3"

	// TextParameters is the constant for the mime type "text/parameters".
	TextParameters = "text/parameters"

	// TextParityfec is the constant for the mime type "text/parityfec".
	TextParityfec = "text/parityfec"

	// TextPlain is the constant for the mime type "text/plain".
	TextPlain = "text/plain"

	// TextProvenanceNotation is the constant for the mime type "text/provenance-notation".
	TextProvenanceNotation = "text/provenance-notation"

	// TextPrsFallensteinRst is the constant for the mime type "text/prs.fallenstein.rst".
	TextPrsFallensteinRst = "text/prs.fallenstein.rst"

	// TextPrsLinesTag is the constant for the mime type "text/prs.lines.tag".
	TextPrsLinesTag = "text/prs.lines.tag"

	// TextPrsPropLogic is the constant for the mime type "text/prs.prop.logic".
	TextPrsPropLogic = "text/prs.prop.logic"

	// TextRaptorfec is the constant for the mime type "text/raptorfec".
	TextRaptorfec = "text/raptorfec"

	// TextRED is the constant for the mime type "text/RED".
	TextRED = "text/RED"

	// TextRfc822Headers is the constant for the mime type "text/rfc822-headers".
	TextRfc822Headers = "text/rfc822-headers"

	// TextRichtext is the constant for the mime type "text/richtext".
	TextRichtext = "text/richtext"

	// TextRtf is the constant for the mime type "text/rtf".
	TextRtf = "text/rtf"

	// TextRtpEncAescm128 is the constant for the mime type "text/rtp-enc-aescm128".
	TextRtpEncAescm128 = "text/rtp-enc-aescm128"

	// TextRtploopback is the constant for the mime type "text/rtploopback".
	TextRtploopback = "text/rtploopback"

	// TextRtx is the constant for the mime type "text/rtx".
	TextRtx = "text/rtx"

	// TextSgml is the constant for the mime type "text/sgml".
	TextSgml = "text/sgml"

	// TextStrings is the constant for the mime type "text/strings".
	TextStrings = "text/strings"

	// TextT140 is the constant for the mime type "text/t140".
	TextT140 = "text/t140"

	// TextTabSeparatedValues is the constant for the mime type "text/tab-separated-values".
	TextTabSeparatedValues = "text/tab-separated-values"

	// TextTroff is the constant for the mime type "text/troff".
	TextTroff = "text/troff"

	// TextTurtle is the constant for the mime type "text/turtle".
	TextTurtle = "text/turtle"

	// TextUlpfec is the constant for the mime type "text/ulpfec".
	TextUlpfec = "text/ulpfec"

	// TextUriList is the constant for the mime type "text/uri-list".
	TextUriList = "text/uri-list"

	// TextVcard is the constant for the mime type "text/vcard".
	TextVcard = "text/vcard"

	// TextVndA is the constant for the mime type "text/vnd.a".
	TextVndA = "text/vnd.a"

	// TextVndAbc is the constant for the mime type "text/vnd.abc".
	TextVndAbc = "text/vnd.abc"

	// TextVndAsciiArt is the constant for the mime type "text/vnd.ascii-art".
	TextVndAsciiArt = "text/vnd.ascii-art"

	// TextVndCurl is the constant for the mime type "text/vnd.curl".
	TextVndCurl = "text/vnd.curl"

	// TextVndDebianCopyright is the constant for the mime type "text/vnd.debian.copyright".
	TextVndDebianCopyright = "text/vnd.debian.copyright"

	// TextVndDMClientScript is the constant for the mime type "text/vnd.DMClientScript".
	TextVndDMClientScript = "text/vnd.DMClientScript"

	// TextVndDvbSubtitle is the constant for the mime type "text/vnd.dvb.subtitle".
	TextVndDvbSubtitle = "text/vnd.dvb.subtitle"

	// TextVndEsmertecThemeDescriptor is the constant for the mime type "text/vnd.esmertec.theme-descriptor".
	TextVndEsmertecThemeDescriptor = "text/vnd.esmertec.theme-descriptor"

	// TextVndFly is the constant for the mime type "text/vnd.fly".
	TextVndFly = "text/vnd.fly"

	// TextVndFmiFlexstor is the constant for the mime type "text/vnd.fmi.flexstor".
	TextVndFmiFlexstor = "text/vnd.fmi.flexstor"

	// TextVndGml is the constant for the mime type "text/vnd.gml".
	TextVndGml = "text/vnd.gml"

	// TextVndGraphviz is the constant for the mime type "text/vnd.graphviz".
	TextVndGraphviz = "text/vnd.graphviz"

	// TextVndHgl is the constant for the mime type "text/vnd.hgl".
	TextVndHgl = "text/vnd.hgl"

	// TextVndIn3D3dml is the constant for the mime type "text/vnd.in3d.3dml".
	TextVndIn3D3dml = "text/vnd.in3d.3dml"

	// TextVndIn3DSpot is the constant for the mime type "text/vnd.in3d.spot".
	TextVndIn3DSpot = "text/vnd.in3d.spot"

	// TextVndIPTCNewsML is the constant for the mime type "text/vnd.IPTC.NewsML".
	TextVndIPTCNewsML = "text/vnd.IPTC.NewsML"

	// TextVndIPTCNITF is the constant for the mime type "text/vnd.IPTC.NITF".
	TextVndIPTCNITF = "text/vnd.IPTC.NITF"

	// TextVndLatexZ is the constant for the mime type "text/vnd.latex-z".
	TextVndLatexZ = "text/vnd.latex-z"

	// TextVndMotorolaReflex is the constant for the mime type "text/vnd.motorola.reflex".
	TextVndMotorolaReflex = "text/vnd.motorola.reflex"

	// TextVndMsMediapackage is the constant for the mime type "text/vnd.ms-mediapackage".
	TextVndMsMediapackage = "text/vnd.ms-mediapackage"

	// TextVndNet2PhoneCommcenterCommand is the constant for the mime type "text/vnd.net2phone.commcenter.command".
	TextVndNet2PhoneCommcenterCommand = "text/vnd.net2phone.commcenter.command"

	// TextVndRadisysMsmlBasicLayout is the constant for the mime type "text/vnd.radisys.msml-basic-layout".
	TextVndRadisysMsmlBasicLayout = "text/vnd.radisys.msml-basic-layout"

	// TextVndSenxWarpscript is the constant for the mime type "text/vnd.senx.warpscript".
	TextVndSenxWarpscript = "text/vnd.senx.warpscript"

	// TextVndSiUricatalogue is the constant for the mime type "text/vnd.si.uricatalogue - OBSOLETED by request".
	TextVndSiUricatalogue = "text/vnd.si.uricatalogue"

	// TextVndSunJ2MeAppDescriptor is the constant for the mime type "text/vnd.sun.j2me.app-descriptor".
	TextVndSunJ2MeAppDescriptor = "text/vnd.sun.j2me.app-descriptor"

	// TextVndSosi is the constant for the mime type "text/vnd.sosi".
	TextVndSosi = "text/vnd.sosi"

	// TextVndTrolltechLinguist is the constant for the mime type "text/vnd.trolltech.linguist".
	TextVndTrolltechLinguist = "text/vnd.trolltech.linguist"

	// TextVndWapSi is the constant for the mime type "text/vnd.wap.si".
	TextVndWapSi = "text/vnd.wap.si"

	// TextVndWapSl is the constant for the mime type "text/vnd.wap.sl".
	TextVndWapSl = "text/vnd.wap.sl"

	// TextVndWapWml is the constant for the mime type "text/vnd.wap.wml".
	TextVndWapWml = "text/vnd.wap.wml"

	// TextVndWapWmlscript is the constant for the mime type "text/vnd.wap.wmlscript".
	TextVndWapWmlscript = "text/vnd.wap.wmlscript"

	// TextXml is the constant for the mime type "text/xml".
	TextXml = "text/xml"

	// TextXmlExternalParsedEntity is the constant for the mime type "text/xml-external-parsed-entity".
	TextXmlExternalParsedEntity = "text/xml-external-parsed-entity"
)

const (
	// Video1dInterleavedParityfec is the constant for the mime type "video/1d-interleaved-parityfec".
	Video1dInterleavedParityfec = "video/1d-interleaved-parityfec"

	// Video3gpp is the constant for the mime type "video/3gpp".
	Video3gpp = "video/3gpp"

	// Video3gpp2 is the constant for the mime type "video/3gpp2".
	Video3gpp2 = "video/3gpp2"

	// Video3gppTt is the constant for the mime type "video/3gpp-tt".
	Video3gppTt = "video/3gpp-tt"

	// VideoBMPEG is the constant for the mime type "video/BMPEG".
	VideoBMPEG = "video/BMPEG"

	// VideoBT656 is the constant for the mime type "video/BT656".
	VideoBT656 = "video/BT656"

	// VideoCelB is the constant for the mime type "video/CelB".
	VideoCelB = "video/CelB"

	// VideoDV is the constant for the mime type "video/DV".
	VideoDV = "video/DV"

	// VideoEncaprtp is the constant for the mime type "video/encaprtp".
	VideoEncaprtp = "video/encaprtp"

	// VideoExample is the constant for the mime type "video/example".
	VideoExample = "video/example"

	// VideoFlexfec is the constant for the mime type "video/flexfec".
	VideoFlexfec = "video/flexfec"

	// VideoH261 is the constant for the mime type "video/H261".
	VideoH261 = "video/H261"

	// VideoH263 is the constant for the mime type "video/H263".
	VideoH263 = "video/H263"

	// VideoH2631998 is the constant for the mime type "video/H263-1998".
	VideoH2631998 = "video/H263-1998"

	// VideoH2632000 is the constant for the mime type "video/H263-2000".
	VideoH2632000 = "video/H263-2000"

	// VideoH264 is the constant for the mime type "video/H264".
	VideoH264 = "video/H264"

	// VideoH264RCDO is the constant for the mime type "video/H264-RCDO".
	VideoH264RCDO = "video/H264-RCDO"

	// VideoH264SVC is the constant for the mime type "video/H264-SVC".
	VideoH264SVC = "video/H264-SVC"

	// VideoH265 is the constant for the mime type "video/H265".
	VideoH265 = "video/H265"

	// VideoIsoSegment is the constant for the mime type "video/iso.segment".
	VideoIsoSegment = "video/iso.segment"

	// VideoJPEG is the constant for the mime type "video/JPEG".
	VideoJPEG = "video/JPEG"

	// VideoJpeg2000 is the constant for the mime type "video/jpeg2000".
	VideoJpeg2000 = "video/jpeg2000"

	// VideoMj2 is the constant for the mime type "video/mj2".
	VideoMj2 = "video/mj2"

	// VideoMP1S is the constant for the mime type "video/MP1S".
	VideoMP1S = "video/MP1S"

	// VideoMP2P is the constant for the mime type "video/MP2P".
	VideoMP2P = "video/MP2P"

	// VideoMP2T is the constant for the mime type "video/MP2T".
	VideoMP2T = "video/MP2T"

	// VideoMp4 is the constant for the mime type "video/mp4".
	VideoMp4 = "video/mp4"

	// VideoMP4VES is the constant for the mime type "video/MP4V-ES".
	VideoMP4VES = "video/MP4V-ES"

	// VideoMPV is the constant for the mime type "video/MPV".
	VideoMPV = "video/MPV"

	// VideoMpeg is the constant for the mime type "video/mpeg".
	VideoMpeg = "video/mpeg"

	// VideoMpeg4Generic is the constant for the mime type "video/mpeg4-generic".
	VideoMpeg4Generic = "video/mpeg4-generic"

	// VideoNv is the constant for the mime type "video/nv".
	VideoNv = "video/nv"

	// VideoOgg is the constant for the mime type "video/ogg".
	VideoOgg = "video/ogg"

	// VideoParityfec is the constant for the mime type "video/parityfec".
	VideoParityfec = "video/parityfec"

	// VideoPointer is the constant for the mime type "video/pointer".
	VideoPointer = "video/pointer"

	// VideoQuicktime is the constant for the mime type "video/quicktime".
	VideoQuicktime = "video/quicktime"

	// VideoRaptorfec is the constant for the mime type "video/raptorfec".
	VideoRaptorfec = "video/raptorfec"

	// VideoRaw is the constant for the mime type "video/raw".
	VideoRaw = "video/raw"

	// VideoRtpEncAescm128 is the constant for the mime type "video/rtp-enc-aescm128".
	VideoRtpEncAescm128 = "video/rtp-enc-aescm128"

	// VideoRtploopback is the constant for the mime type "video/rtploopback".
	VideoRtploopback = "video/rtploopback"

	// VideoRtx is the constant for the mime type "video/rtx".
	VideoRtx = "video/rtx"

	// VideoSmpte291 is the constant for the mime type "video/smpte291".
	VideoSmpte291 = "video/smpte291"

	// VideoSMPTE292M is the constant for the mime type "video/SMPTE292M".
	VideoSMPTE292M = "video/SMPTE292M"

	// VideoUlpfec is the constant for the mime type "video/ulpfec".
	VideoUlpfec = "video/ulpfec"

	// VideoVc1 is the constant for the mime type "video/vc1".
	VideoVc1 = "video/vc1"

	// VideoVc2 is the constant for the mime type "video/vc2".
	VideoVc2 = "video/vc2"

	// VideoVndCCTV is the constant for the mime type "video/vnd.CCTV".
	VideoVndCCTV = "video/vnd.CCTV"

	// VideoVndDeceHd is the constant for the mime type "video/vnd.dece.hd".
	VideoVndDeceHd = "video/vnd.dece.hd"

	// VideoVndDeceMobile is the constant for the mime type "video/vnd.dece.mobile".
	VideoVndDeceMobile = "video/vnd.dece.mobile"

	// VideoVndDeceMp4 is the constant for the mime type "video/vnd.dece.mp4".
	VideoVndDeceMp4 = "video/vnd.dece.mp4"

	// VideoVndDecePd is the constant for the mime type "video/vnd.dece.pd".
	VideoVndDecePd = "video/vnd.dece.pd"

	// VideoVndDeceSd is the constant for the mime type "video/vnd.dece.sd".
	VideoVndDeceSd = "video/vnd.dece.sd"

	// VideoVndDeceVideo is the constant for the mime type "video/vnd.dece.video".
	VideoVndDeceVideo = "video/vnd.dece.video"

	// VideoVndDirectvMpeg is the constant for the mime type "video/vnd.directv.mpeg".
	VideoVndDirectvMpeg = "video/vnd.directv.mpeg"

	// VideoVndDirectvMpegTts is the constant for the mime type "video/vnd.directv.mpeg-tts".
	VideoVndDirectvMpegTts = "video/vnd.directv.mpeg-tts"

	// VideoVndDlnaMpegTts is the constant for the mime type "video/vnd.dlna.mpeg-tts".
	VideoVndDlnaMpegTts = "video/vnd.dlna.mpeg-tts"

	// VideoVndDvbFile is the constant for the mime type "video/vnd.dvb.file".
	VideoVndDvbFile = "video/vnd.dvb.file"

	// VideoVndFvt is the constant for the mime type "video/vnd.fvt".
	VideoVndFvt = "video/vnd.fvt"

	// VideoVndHnsVideo is the constant for the mime type "video/vnd.hns.video".
	VideoVndHnsVideo = "video/vnd.hns.video"

	// VideoVndIptvforum1dparityfec1010 is the constant for the mime type "video/vnd.iptvforum.1dparityfec-1010".
	VideoVndIptvforum1dparityfec1010 = "video/vnd.iptvforum.1dparityfec-1010"

	// VideoVndIptvforum1dparityfec2005 is the constant for the mime type "video/vnd.iptvforum.1dparityfec-2005".
	VideoVndIptvforum1dparityfec2005 = "video/vnd.iptvforum.1dparityfec-2005"

	// VideoVndIptvforum2dparityfec1010 is the constant for the mime type "video/vnd.iptvforum.2dparityfec-1010".
	VideoVndIptvforum2dparityfec1010 = "video/vnd.iptvforum.2dparityfec-1010"

	// VideoVndIptvforum2dparityfec2005 is the constant for the mime type "video/vnd.iptvforum.2dparityfec-2005".
	VideoVndIptvforum2dparityfec2005 = "video/vnd.iptvforum.2dparityfec-2005"

	// VideoVndIptvforumTtsavc is the constant for the mime type "video/vnd.iptvforum.ttsavc".
	VideoVndIptvforumTtsavc = "video/vnd.iptvforum.ttsavc"

	// VideoVndIptvforumTtsmpeg2 is the constant for the mime type "video/vnd.iptvforum.ttsmpeg2".
	VideoVndIptvforumTtsmpeg2 = "video/vnd.iptvforum.ttsmpeg2"

	// VideoVndMotorolaVideo is the constant for the mime type "video/vnd.motorola.video".
	VideoVndMotorolaVideo = "video/vnd.motorola.video"

	// VideoVndMotorolaVideop is the constant for the mime type "video/vnd.motorola.videop".
	VideoVndMotorolaVideop = "video/vnd.motorola.videop"

	// VideoVndMpegurl is the constant for the mime type "video/vnd.mpegurl".
	VideoVndMpegurl = "video/vnd.mpegurl"

	// VideoVndMsPlayreadyMediaPyv is the constant for the mime type "video/vnd.ms-playready.media.pyv".
	VideoVndMsPlayreadyMediaPyv = "video/vnd.ms-playready.media.pyv"

	// VideoVndNokiaInterleavedMultimedia is the constant for the mime type "video/vnd.nokia.interleaved-multimedia".
	VideoVndNokiaInterleavedMultimedia = "video/vnd.nokia.interleaved-multimedia"

	// VideoVndNokiaMp4Vr is the constant for the mime type "video/vnd.nokia.mp4vr".
	VideoVndNokiaMp4Vr = "video/vnd.nokia.mp4vr"

	// VideoVndNokiaVideovoip is the constant for the mime type "video/vnd.nokia.videovoip".
	VideoVndNokiaVideovoip = "video/vnd.nokia.videovoip"

	// VideoVndObjectvideo is the constant for the mime type "video/vnd.objectvideo".
	VideoVndObjectvideo = "video/vnd.objectvideo"

	// VideoVndRadgamettoolsBink is the constant for the mime type "video/vnd.radgamettools.bink".
	VideoVndRadgamettoolsBink = "video/vnd.radgamettools.bink"

	// VideoVndRadgamettoolsSmacker is the constant for the mime type "video/vnd.radgamettools.smacker".
	VideoVndRadgamettoolsSmacker = "video/vnd.radgamettools.smacker"

	// VideoVndSealedMpeg1 is the constant for the mime type "video/vnd.sealed.mpeg1".
	VideoVndSealedMpeg1 = "video/vnd.sealed.mpeg1"

	// VideoVndSealedMpeg4 is the constant for the mime type "video/vnd.sealed.mpeg4".
	VideoVndSealedMpeg4 = "video/vnd.sealed.mpeg4"

	// VideoVndSealedSwf is the constant for the mime type "video/vnd.sealed.swf".
	VideoVndSealedSwf = "video/vnd.sealed.swf"

	// VideoVndSealedmediaSoftsealMov is the constant for the mime type "video/vnd.sealedmedia.softseal.mov".
	VideoVndSealedmediaSoftsealMov = "video/vnd.sealedmedia.softseal.mov"

	// VideoVndUvvuMp4 is the constant for the mime type "video/vnd.uvvu.mp4".
	VideoVndUvvuMp4 = "video/vnd.uvvu.mp4"

	// VideoVndYoutubeYt is the constant for the mime type "video/vnd.youtube.yt".
	VideoVndYoutubeYt = "video/vnd.youtube.yt"

	// VideoVndVivo is the constant for the mime type "video/vnd.vivo".
	VideoVndVivo = "video/vnd.vivo"

	// VideoVP8 is the constant for the mime type "video/VP8".
	VideoVP8 = "video/VP8"
)
