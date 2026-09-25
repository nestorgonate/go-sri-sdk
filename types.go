package gosrisdk

type ConsultaRuc struct {
	Contribuyentes []Contribuyente `json:"contribuyentes"`
	Encuensta      Encuesta        `json:"encuesta"`
}

type Contribuyente struct {
	NumeroRuc                      string                        `json:"numeroRuc"`
	RazonSocial                    string                        `json:"razonSocial"`
	EstadoContribuyenteRuc         string                        `json:"estadoContribuyenteRuc"`
	ActividadEconomicaPrincipal    string                        `json:"actividadEconomicaPrincipal"`
	TipoContribuyente              string                        `json:"tipoContribuyente"`
	Regimen                        string                        `json:"regimen"`
	Categoria                      *string                       `json:"categoria"`
	ObligadoLlevarContabilidad     SriBool                       `json:"obligadoLlevarContabilidad"`
	AgenteRetencion                SriBool                       `json:"agenteRetencion"`
	ContribuyenteEspecial          SriBool                       `json:"contribuyenteEspecial"`
	InformacionFechasContribuyente InformacionFechaContribuyente `json:"informacionFechasContribuyente"`
	RepresentantesLegales          []RepresentanteLegal          `json:"representantesLegales"`
	MotivoCancelacionSuspension    *MotivoCancelacionSuspension  `json:"motivoCancelacionSuspension"`
	ContribuyenteFantasma          SriBool                       `json:"contribuyenteFantasma"`
	TransaccionesInexistente       SriBool                       `json:"transaccionesInexistente"`
}

type InformacionFechaContribuyente struct {
	FechaInicioActividades   SriDate `json:"fechaInicioActividades"`
	FechaCese                SriDate `json:"fechaCese"`
	FechaReinicioActividades SriDate `json:"fechaReinicioActividades"`
	FechaActualizacion       SriDate `json:"fechaActualizacion"`
}

type RepresentanteLegal struct {
	Identificacion string `json:"identificacion"`
	Nombre         string `json:"nombre"`
}

// Consultar un RUC suspendido para saber los datos
type MotivoCancelacionSuspension struct {
}

type Encuesta struct {
	Habilitada bool   `json:"habilitada"`
	Url        string `json:"url"`
}

type Establecimiento struct {
	ID                      string  `json:"numeroEstablecimiento"`
	NombreFantasiaComercial string  `json:"nombreFantasiaComercial"`
	TipoEstablecimiento     string  `json:"tipoEstablecimiento"`
	DireccionCompleta       string  `json:"direccionCompleta"`
	Estado                  string  `json:"estado"`
	Matriz                  SriBool `json:"matriz"`
}
